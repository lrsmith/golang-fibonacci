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

	jsonData, _ := json.Marshal(status)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status.HTTPStatus)

	io.WriteString(w, string(jsonData))

}
