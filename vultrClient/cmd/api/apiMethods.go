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
