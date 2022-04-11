// This package contains the types to handle the responses of the Vultr API
package main

import (
	"time"
)

// RequestCreateInstance has the info necessary to request the creation of a
// new instance. If a value has the omitempty tag, then it is not strictly
// required to request the creation of a new instance.
type RequestCreateInstance struct {
	OS_ID    int    `json:"os_id"`
	Label    string `json:"label,omitempty"`
	Hostname string `json:"hostname,omitempty"`
	Region   string `json:"region"`
	Plan     string `json:"plan"`
	// Backups [enabled|disabled] automatic backups for the new instance.
	// Enabling automatic backups makes the instance more expensive..
	Backups string `json:"backups,omitempty"`
	// SSHKeys is a slice with the IDs of the SSH keys that should be used per
	// default when initializing the newly requested instance
	SSHKeys []string `json:"sshkey_id,omitempty"`
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
	OSID             int       `json:"os_id"`
	AppID            int       `json:"app_id"`
	ImageID          string    `json:"image_id"`
	FirewallGroupID  string    `json:"firewall_group_id"`
	Features         []string  `json:"features"`
	DefaultPassword  string    `json:"default_password"`
}

// ResponseLiveInstance has the info received from response after requesting
// the information from a live instance. The JSON response encapsulates the
// actual instance's data inside a JSON object first.
type ResponseLiveInstance struct {
	Instance *LiveInstance `json:"instance"`
}

// LiveInstance is the data structure that describes a live instance
type LiveInstance struct {
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
	OSID             int       `json:"os_id"`
	AppID            int       `json:"app_id"`
	ImageID          string    `json:"image_id"`
	FirewallGroupID  string    `json:"firewall_group_id"`
	Features         []string  `json:"features"`
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
