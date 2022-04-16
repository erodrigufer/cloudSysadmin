package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// createInstance, create an Instance with the given plan in the specified
// regions, e.g. region="ewr" (New Jersey), plan="vc2-1c-1gb".
// Additionally, use a label and a hostname for the new instance, all this
// information is contained in the fields of the newInstance parameter.
func (app *application) createInstance(newInstance *RequestCreateInstance) (*CreatedInstance, error) {
	// Create a buffer with Read/Write methods implemented
	buf := new(bytes.Buffer)
	// Encode the data to JSON
	err := json.NewEncoder(buf).Encode(newInstance)
	if err != nil {
		return nil, err
	}

	URLInstancesVultrAPI := "https://api.vultr.com/v2/instances"
	req, err := http.NewRequest("POST", URLInstancesVultrAPI, buf)
	if err != nil {
		return nil, err
	}

	app.addAuthToken(req)

	// Send request to Vultr API
	resp, err := app.client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	// if everything went well Vultr API responds with 202 (Status Accepted)
	if err = checkResponseAPI(resp, http.StatusAccepted); err != nil {
		return nil, err
	}

	responseJSON := new(ResponseCreateInstance)
	// decode response into JSON
	err = json.NewDecoder(resp.Body).Decode(responseJSON)
	if err != nil {
		return nil, err
	}
	// since the json response encapsulates the actual instance information,
	// inside an instance object, we have to return only this object
	return responseJSON.Instance, nil
}

// getInstance, fetch information of instance already deployed in the Vultr
// cloud. A pointer to a createdInstance as a parameter has all the information
// required to perform API call.
func (app *application) getInstance(instance *CreatedInstance) (*LiveInstance, error) {
	// Create/format string with URL for running instance
	URLGetInstanceVultrAPI := fmt.Sprintf("https://api.vultr.com/v2/instances/%s", instance.ID)
	req, err := http.NewRequest("GET", URLGetInstanceVultrAPI, nil)
	if err != nil {
		return nil, err
	}

	app.addAuthToken(req)

	// Send request to Vultr API
	resp, err := app.client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	// if everything went well Vultr API responds with 200 (OK)
	if err = checkResponseAPI(resp, http.StatusOK); err != nil {
		return nil, err
	}

	responseJSON := new(ResponseLiveInstance)
	// decode response into JSON
	err = json.NewDecoder(resp.Body).Decode(responseJSON)
	if err != nil {
		return nil, err
	}
	// since the json response encapsulates the actual instance information,
	// inside an instance object, we have to return only this object
	return responseJSON.Instance, nil

}

// deleteInstance, deletes an actively running instance. The input parameter
// is a pointer to an instance of type 'LiveInstance'
func (app *application) deleteInstance(instance *LiveInstance) error {
	// Create/format string with URL for running instance
	URLDeleteInstanceVultrAPI := fmt.Sprintf("https://api.vultr.com/v2/instances/%s", instance.ID)
	req, err := http.NewRequest("DELETE", URLDeleteInstanceVultrAPI, nil)
	if err != nil {
		return err
	}

	app.addAuthToken(req)

	// Send request to Vultr API
	resp, err := app.client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	// if everything went well Vultr API responds with 204 (No Content)
	// more information regarding this behaviour can be found here:
	// https://datatracker.ietf.org/doc/draft-ietf-httpbis-semantics/
	// If a DELETE method is successfully applied, the origin server SHOULD send
	//	*  a 202 (Accepted) status code if the action will likely succeed but
	//	has not yet been enacted,
	//
	//	*  a 204 (No Content) status code if the action has been enacted and
	//	no further information is to be supplied, or
	//
	//	*  a 200 (OK) status code if the action has been enacted and the
	//	response message includes a representation describing the status.
	// 	[...] The _204 (No Content)_ status code indicates that the server has
	// 	successfully fulfilled the request and that there is no additional
	//	content to send in the response content.
	if err = checkResponseAPI(resp, http.StatusNoContent); err != nil {
		return err
	}

	return nil

}

// listSSHKeys, list all SSH Keys registered to a particular Vultr account
func (app *application) listSSHKeys() ([]SSHKey, error) {
	URLSSHKeysVultrAPI := "https://api.vultr.com/v2/ssh-keys"
	req, err := http.NewRequest("GET", URLSSHKeysVultrAPI, nil)
	app.addAuthToken(req)

	// Send request to Vultr API
	resp, err := app.client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	// if everything went well VultrAPI respondes with 200 (OK)
	if err = checkResponseAPI(resp, http.StatusOK); err != nil {
		return nil, err
	}

	responseJson := new(ResponseListSSHKeys)
	err = json.NewDecoder(resp.Body).Decode(responseJson)
	if err != nil {
		return nil, err
	}

	return responseJson.SSHKeys, nil
}
