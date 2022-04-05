package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

//type InstanceCreated struct {
//ID string "id": "4f0f12e5-1f84-404f-aa84-85f431ea5ec2",
//"os": "CentOS 8 Stream",
//"ram": 1024,
//"disk": 0,
//"main_ip": "0.0.0.0",
//"vcpu_count": 1,
//"region": "ewr",
//"plan": "vc2-1c-1gb",
//"date_created": "2021-09-14T13:22:20+00:00",
//"status": "pending",
//"allowed_bandwidth": 1000,
//"netmask_v4": "",
//"gateway_v4": "0.0.0.0",
//"power_status": "running",
//"server_status": "none",
//"v6_network": "",
//"v6_main_ip": "",
//"v6_network_size": 0,
//"label": "",
//"internal_ip": "",
//"kvm": "",
//"hostname": "my_hostname",
//"tag": "",
//"os_id": 401,
//"app_id": 0,
//"image_id": "",
//"firewall_group_id": "",
//"features": [],
//Password string "default_password": "v5{Fkvb#2ycPGwHs"
//}

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

	resp.Write(os.Stdout)
	body, err := io.ReadAll(resp.Body)

	responseDecoded := new(InstanceCreated)
	// decode response into JSON
	err = json.Unmarshal(body, responseDecoded)
}
