package main

import (
	"errors"
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

// Check if the response from the API is correct, if not return error
// The input parameter is the expected correct HTTP response
func checkResponseAPI(resp *http.Response, correctResponse int) (err error) {
	// if everything went well Vultr API responds with "correctResponse"
	if resp.StatusCode != correctResponse {
		errorMessage := fmt.Sprintf("request failed. Response status: %s", resp.Status)
		err = errors.New(errorMessage)
		return
	}
	return nil
}
