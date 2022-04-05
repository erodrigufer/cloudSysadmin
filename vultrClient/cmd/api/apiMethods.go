package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// URL of Vultr API
const VultrAPI = "https://api.vultr.com/v2/instances"

type Instance struct {
	// omitempty= if value is not present, omit at encoding
	ID       string `json:"id,omitempty"`
	OS_id    int    `json:"os_id"`
	Label    string `json:"label,omitempty"`
	Hostname string `json:"hostname,omitempty"`
	Region   string `json:"region"`
	Plan     string `json:"plan"`
	Backups  string `json:"backups"` // disabled (no backups)
	// with backups is more expensive
}

// Create an Instance with the given plan in the specified regions
// e.g. region="ewr" (New Jersey), plan="vc2-1c-1gb".
// Additionally, use a label and a hostname for the new instance
func (app *application) createInstance(newInstance *Instance) {
	// Validate input parameters, plans
	//fmt.Println(newInstance.Region)
	//app.infoLog.Println("%s", app.cfg.tokenAPI)

	client := new(http.Client)

	// Create a buffer with Read/Write methods implemented
	buf := new(bytes.Buffer)
	// Encode the data to JSON
	err := json.NewEncoder(buf).Encode(newInstance)
	if err != nil {
		app.errorLog.Fatal("json encoding failed")
	}

	req, err := http.NewRequest("POST", VultrAPI, buf)
	if err != nil {
		app.errorLog.Fatal("request creation failed")
	}

	// Format token value for header. Add 'Bearer' before token
	tokenValue := fmt.Sprint("Bearer ", app.cfg.tokenAPI)
	req.Header.Add("Authorization", tokenValue)
	req.Header.Add("Content-Type", "application/json")

	// Print req to stdout
	// req.Write(os.Stdout)

	// Send request to Vultr API
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		app.errorLog.Println("client send request failed")
	}
	// if everything went well Vultr API responds with 202 (Status Accepted)
	if resp.Status != http.StatusAccepted {
		app.errorLog.Println("server did not accept request. Response status: ", resp.Status)
	}
	//	body, err := io.ReadAll(resp.Body)
	resp.Write(os.Stdout)

}
