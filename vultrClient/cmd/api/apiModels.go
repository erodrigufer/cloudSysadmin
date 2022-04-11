// This package contains the types to handle the responses of the Vultr API
package main

import (
	"time"
)

// RequestCreateInstance has the info necessary to request the creation of a
// new instance
type RequestCreateInstance struct {
	ID       string   `json:"id,omitempty"` // omitempty= if value is not present, omit at encoding
	OS_id    int      `json:"os_id"`
	Label    string   `json:"label,omitempty"`
	Hostname string   `json:"hostname,omitempty"`
	Region   string   `json:"region"`
	Plan     string   `json:"plan"`
	Backups  string   `json:"backups"` // disabled (no backups)
	SSHKeys  []string `json:"sshkey_id,omitempty"`
	// with backups is more expensive
}

// ResponseCreateInstance has the info received from response after requesting
// to create an instance
type ResponseCreateInstance struct {
	Instance *CreatedInstance `json:"instance"`
}

// CreatedInstance is the data structure that describes a created instances
type CreatedInstance struct {
	ID        string `json:"id"`
	OS        string `json:"os"`
	RAM       int    `json:"ram"`
	Disk      int    `json:"disk"`
	MainIP    string `json:"main_ip"`
	VCPUCount int    `json:"vcpu_count"`
	// Region is the region where the instance is being deployed
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

// SSHKey is the data structure with the fields that describe a single SSH key
type SSHKey struct {
	ID           string    `json:"id"`
	CreationDate time.Time `json:"date_created"`
	// Name is the name given to a single SSH key
	Name string `json:"name"`
	// SSH_Pub_Key is the actual public ssh key
	SSH_Pub_Key string `json:"ssh_key"`
}

// ResponseListSSHKeys has the information received after requesting a list with
// the stored SSH keys in the system
type ResponseListSSHKeys struct {
	SSHKeys []SSHKey `json:"ssh_keys"` // slice of ssh keys
	Meta    struct {
		TotalKeys int `json:"total"` // total amount of ssh keys
		Links     struct {
			Next string `json:"next"`
			Prev string `json:prev"`
		}
	}
}
