// Eduardo Rodriguez @erodrigufer (c) 2022

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

// application type used for dependency injection / avoid using globals
type application struct {
	errorLog *log.Logger   // error log handler
	infoLog  *log.Logger   // info log handler
	cfg      *configValues // config values (e.g. from flags)
	client   *http.Client
}

// store all flag-parseable config values in this struct
type configValues struct {
	tokenAPI string // API token used to communicate with Vultr API
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
	flag.StringVar(&cfg.tokenAPI, "tokenAPI", "-", "Personal Access Token to interact with Vultr API")
	flag.Parse()
	if cfg.tokenAPI == "-" {
		app.errorLog.Fatal("missing API token")
	}

	app.cfg = cfg

	// Use a single http.Client for all interactions with the API, as stated in
	// the official Go documentation "The Client's Transport typically has
	// internal state (cached TCP connections), so Clients should be reused
	// instead of created as needed. Clients are safe for concurrent use by
	// multiple goroutines."
	app.client = new(http.Client)

	newInstance := &Instance{
		OS_id:   447,   // FreeBSD-13
		Region:  "fra", // New Jersey (ewr) Frankfurt (fra)
		Backups: "disabled",
		// Enabling backups makes the VM more expensive
		Plan: "vc2-1c-1gb",
	}

	createdInstance, err := app.createInstance(newInstance)
	if err != nil {
		app.errorLog.Println(err)
		return
	}
	fmt.Printf("%+v\n", createdInstance)
	fmt.Println("New instance's ID: ", createdInstance.ID)

	fmt.Println(app.listSSHKeys)
}
