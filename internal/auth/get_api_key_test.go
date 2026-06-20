package auth

import (
	"net/http"
	"testing"
)

var token = "dXNlcm5hbWU6cGFzc3dvcmQ="

func TestGetAPIKey(t *testing.T) {
	h := http.Header{}
	bearerToken := "ApiKey " + token
	h.Add("Authorization", bearerToken)
	apiKey, err := GetAPIKey(h)
	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}

	if apiKey != token {
		t.Errorf("Expected: %s but got %s", token, apiKey)
	}
}

func TestGetAPIKeyError(t *testing.T) {
	h := http.Header{}
	bearerToken := "ApiKey " + token
	h.Add("Authorization", bearerToken)
	_, err := GetAPIKey(h)

	if err != nil && err != ErrNoAuthHeaderIncluded {
		t.Errorf("Expected error to be %v but got %v", ErrNoAuthHeaderIncluded, err)
	}
}
