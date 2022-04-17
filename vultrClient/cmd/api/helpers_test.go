package main

import (
	"net/http"
	"testing"
)

func TestCheckResponseAPI(t *testing.T) {
	respOK := new(http.Response)
	respOK.StatusCode = 200 // 200 OK

	if err := checkResponseAPI(respOK, 200); err != nil {
		t.Errorf("checkResponseAPI did not recognize 200 OK")
	}

	if err := checkResponseAPI(respOK, 404); err == nil {
		t.Errorf("checkResponseAPI did not recognize different status codes")
	}

}
