// Eduardo Rodriguez @erodrigufer (c) 2022

package main

import (
	"flag"
	"log"
	"os"
)

// application type used for dependency injection / avoid using globals
type application struct {
	errorLog *log.Logger   // error log handler
	infoLog  *log.Logger   // info log handler
	cfg      *configValues // config values (e.g. from flags)
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

	// Config values are valid
	app.cfg = cfg

	newInstance := &Instance{
		OS_id:   447,   // FreeBSD-13
		Region:  "ewr", // New Jersey
		Backups: "disabled",
		// Enabling backups makes the VM more expensive
		Plan: "vc2-1c-1gb",
	}

	app.createInstance(newInstance)
}
