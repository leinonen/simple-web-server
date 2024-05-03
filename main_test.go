package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMainHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	mainHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expectedSubstring := "<h1>Hello, World!</h1>"
	responseBody := rr.Body.String()

	// Check if the response body contains the expected substring
	if !strings.Contains(responseBody, expectedSubstring) {
		t.Errorf("handler returned unexpected body: got %v which does not contain %v", responseBody, expectedSubstring)
	}
}
