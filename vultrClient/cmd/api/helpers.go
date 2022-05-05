package main

import (
	"fmt"
	"net/http"
)

// addAuthToken, adds the necessary headers to get authenticated by the
// Vultr API.
func (app *application) addAuthToken(req *http.Request) {
	// Format token value for header. Add 'Bearer' before token
	tokenValue := fmt.Sprint("Bearer ", app.cfg.tokenAPI)
	req.Header.Add("Authorization", tokenValue)
	req.Header.Add("Content-Type", "application/json")
}

// checkResponseAPI, checks if the response from the API is correct, if not
// it returns error.
// The input parameter is the expected correct HTTP response (int) and the
// http.Response being checked.
func checkResponseAPI(resp *http.Response, correctResponse int) (err error) {
	// if everything went well Vultr API responds with "correctResponse" (int)
	if resp.StatusCode != correctResponse {
		if resp.StatusCode == 401 {
			err = fmt.Errorf("request failed. Response status: %s, expected %d %s. A 401 error response normally takes place, because the IP from which the request was done is not green-listed to communicate with the Vultr API. Check under 'Account/API/Access Control' in the Vultr control panel, if you need to enable access to the API for further IPv4 and IPv6 addresses to perform your next request.", resp.Status, correctResponse, http.StatusText(correctResponse))
			return
		}
		err = fmt.Errorf("request failed. Response status: %s, expected %d", resp.Status, correctResponse)
		return
	}
	return nil
}
