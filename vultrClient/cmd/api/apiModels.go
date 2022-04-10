// This package contains the types to handle the responses of the Vultr API
package main

import (
	"time"
)

// Info to request the creation of a new instance
type RequestCreateInstance struct {
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

// Info received from response after creating instance
type ResponseCreateInstance struct {
	Instance *CreatedInstance `json:"instance"`
}

type CreatedInstance struct {
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

type SSHKey struct {
	ID           string    `json:"id"`
	CreationDate time.Time `json:"date_created"`
	Name         string    `json:"name"`
	SSH_Pub_Key  string    `json:"ssh_key"`
}

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
