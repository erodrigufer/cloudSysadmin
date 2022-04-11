// Eduardo Rodriguez @erodrigufer (c) 2022

package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// application type used for dependency injection / avoid using globals
type application struct {
	// errorLog, error log handler
	errorLog *log.Logger
	// infoLog, info log handler
	infoLog *log.Logger
	// cfg, config values (e.g. from flags)
	cfg *configValues
	// client, long-living http client used to communicate with API
	client *http.Client
}

// store all flag-parseable config values in this struct
type configValues struct {
	// tokenAPI, API token used to communicate with Vultr API
	tokenAPI string
	// sshkey, ssh key which will be automatically initialized in new instance
	sshkey string
	// hostname, hostname for new instance
	hostname string
	// label, label for new instance
	label string
	// region, region where to deplay VM
	region string
}

func main() {
	// Create a logger for INFO messages, the prefix "INFO" and a tab will be
	// displayed before each log message. The flags Ldate and Ltime provide the
	// local date and time
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// Create an ERROR messages logger, additionally use the Lshortfile flag to
	// display the file's name and line number for the error
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		infoLog:  infoLog,
		errorLog: errorLog,
	}

	cfg := new(configValues)
	flag.StringVar(&cfg.tokenAPI, "tokenAPI", "", "Personal Access Token to interact with Vultr API")
	flag.StringVar(&cfg.sshkey, "sshKey", "", "SSH key to initialize per default in new instance")
	flag.StringVar(&cfg.hostname, "hostname", "", "Hostname for new instance")
	flag.StringVar(&cfg.label, "label", "", "Label for new instance")
	flag.StringVar(&cfg.region, "region", "fra", "Region where to deply VM")
	flag.Parse()
	if cfg.tokenAPI == "" {
		app.errorLog.Fatal("missing API token")
	}

	app.cfg = cfg

	// Use a single http.Client for all interactions with the API, as stated in
	// the official Go documentation "The Client's Transport typically has
	// internal state (cached TCP connections), so Clients should be reused
	// instead of created as needed. Clients are safe for concurrent use by
	// multiple goroutines."
	app.client = new(http.Client)

	newInstance := &RequestCreateInstance{
		OS_ID:   447,            // FreeBSD-13
		Region:  app.cfg.region, // New Jersey (ewr) Frankfurt (fra)
		Backups: "disabled",
		// Enabling backups makes the VM more expensive
		Plan:     "vc2-1c-1gb",
		Hostname: app.cfg.hostname,
		Label:    app.cfg.label,
	}

	// Append ssh key parsed from flags, if no ssh key was parsed an empty
	// string is appended
	newInstance.SSHKeys = append(newInstance.SSHKeys, app.cfg.sshkey)

	createdInstance, err := app.createInstance(newInstance)
	if err != nil {
		app.errorLog.Println(err)
		return
	}
	app.infoLog.Printf("new instance [ID: %s] created.", createdInstance.ID)

	//fmt.Printf("%+v\n", createdInstance)
	//fmt.Println("New instance's ID: ", createdInstance.ID)

	//keys, err := app.listSSHKeys()
	//if err != nil {
	//	app.errorLog.Println(err)
	//	return
	//}
	//fmt.Printf("SSH Keys: %+v\n", keys)
}
