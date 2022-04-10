package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Create an Instance with the given plan in the specified regions,
// e.g. region="ewr" (New Jersey), plan="vc2-1c-1gb".
// Additionally, use a label and a hostname for the new instance
func (app *application) createInstance(newInstance *Instance) (*InstanceCreated, error) {
	// Create a buffer with Read/Write methods implemented
	buf := new(bytes.Buffer)
	// Encode the data to JSON
	err := json.NewEncoder(buf).Encode(newInstance)
	if err != nil {
		return nil, err
	}

	instancesVultrAPI := "https://api.vultr.com/v2/instances"
	req, err := http.NewRequest("POST", instancesVultrAPI, buf)
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

	createdInstance := new(InstanceHack)
	// decode response into JSON
	err = json.NewDecoder(resp.Body).Decode(createdInstance)
	if err != nil {
		return nil, err
	}
	return createdInstance.Instance, nil
}

// List all SSH Keys registered to a particular Vultr account
func (app *application) listSSHKeys() ([]SSHKey, error) {
	sshkeysVultrAPI := "https://api.vultr.com/v2/ssh-keys"
	req, err := http.NewRequest("GET", sshkeysVultrAPI, nil)
	app.addAuthToken(req)

	// Send requesto to Vultr API
	resp, err := app.client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	// if everything went well VultrAPI respondes with 200 (OK)
	if err = checkResponseAPI(resp, http.StatusOK); err != nil {
		return nil, err
	}

	responseJson := new(ResponseSSHKeys)
	err = json.NewDecoder(resp.Body).Decode(responseJson)
	if err != nil {
		return nil, err
	}

	return responseJson.SSHKeys, nil
}
