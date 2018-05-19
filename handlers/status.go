package handlers

import (
	"encoding/json"
	"io"
	"net/http"
)

type statusStruct struct {
	HTTPStatus int    `json:"httpstatus"`
	Status     string `json:"status"`
}

// Initializer for statusStruct
// By default fail closed and assume 'dead'
func newstatusStruct() statusStruct {
	return statusStruct{http.StatusInternalServerError, "dead"}
}

// Status : Handler that returns status/health of rest API. F
// Status should return : healthy, sick, dead
func Status(w http.ResponseWriter, r *http.Request) {

	status := newstatusStruct()

	// Normally would have some code to verify services and then adjust status
	status.HTTPStatus = http.StatusOK
	status.Status = "healthy"

	// Marshal struct into JSON
	jsonData, _ := json.Marshal(status)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status.HTTPStatus)

	// Convert JSON from []bytes to strings and send it back
	io.WriteString(w, string(jsonData))

}
