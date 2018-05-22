package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test requesting the status URI without authentication
func TestUnauthenticatedRequest(t *testing.T) {

	appConfigs := getConfigs()
	router := NewRouter(appConfigs)

	req, _ := http.NewRequest("GET", "/status", nil)

	requestRecorder := httptest.NewRecorder()
	router.ServeHTTP(requestRecorder, req)

	if requestRecorder.Code != http.StatusForbidden {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusForbidden, requestRecorder.Code)
	}
	if requestRecorder.Body.String() != "{\"httpstatus\":403,\"errmsg\":\"Authentication Failure.\"}"+"\n" {
		t.Errorf("Un-expected response %s.", requestRecorder.Body)
	}
}
