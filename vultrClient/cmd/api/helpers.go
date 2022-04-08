package main

import (
	"fmt"
	"net/http"
)

// Add the necessary headers to get authenticated by the Vultr API
func (app *application) addAuthToken(req *http.Request) {
	// Format token value for header. Add 'Bearer' before token
	tokenValue := fmt.Sprint("Bearer ", app.cfg.tokenAPI)
	req.Header.Add("Authorization", tokenValue)
	req.Header.Add("Content-Type", "application/json")

}
