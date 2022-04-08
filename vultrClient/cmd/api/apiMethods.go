package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// URL of Vultr API v2
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

type InstanceHack struct {
	Instance InstanceCreated `json:"instance"`
}

// type info received from response after creating instance
type InstanceCreated struct {
	ID               string    `json:"id"`
	OS               string    `json:"os"`
	RAM              int       `json:"ram"`
	Disk             int       `json:"disk"`
	MainIP           string    `json:"main_ip"`
	VCPUCount        int       `json:"vcpu_count"`
	Region           string    `json:"region"`
	Plan             string    `json:"plan"`
	CreationDate     time.Time `json:"date_created"`
	Status           string    `json:"status"`
	AllowedBandwidth int       `json:"allowed_bandwidth"`
	NetmaskV4        string    `json:"netmask_v4"`
	GatewayV4        string    `json:"gateway_v4"`
	PowerStatus      string    `json:"power_status"`
	ServerStatus     string    `json:"server_status"`
	V6Network        string    `json:"v6_network"`
	V6MainIP         string    `json:"v6_main_ip"`
	V6NetworkSize    int       `json:"v6_network_size"`
	Label            string    `json:"label"`
	InternalIP       string    `json:"internal_ip"`
	KVM              string    `json:"kvm"`
	Hostname         string    `json:"hostname"`
	Tag              string    `json:"tag"`
	OS_id            int       `json:"os_id"`
	App_id           int       `json:"app_id"`
	Image_id         string    `json:"image_id"`
	FirewallGroupID  string    `json:"firewall_group_id"`
	Features         []string  `json:"features"`
	DefaultPassword  string    `json:"default_password"`
}

// Create an Instance with the given plan in the specified regions,
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
	if resp.StatusCode != http.StatusAccepted {
		app.errorLog.Println("server did not accept request. Response status: ", resp.Status)
		return
	}

	//resp.Write(os.Stdout)
	//body, err := io.ReadAll(resp.Body)

	responseDecoded := new(InstanceHack)
	// decode response into JSON
	err = json.NewDecoder(resp.Body).Decode(responseDecoded)
	if err != nil {
		app.errorLog.Println("error decoding json response")
		return
	}
	fmt.Printf("%+v\n", responseDecoded)
	fmt.Println("New instance's ID: ", responseDecoded.Instance.ID)
}
