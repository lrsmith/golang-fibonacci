package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type fibSeqResult struct {
	HTTPStatus int      `json:"httpstatus"`
	Sequence   []uint64 `json:"sequence"`
	ErrorMsg   string   `json:"errormsg"`
}

func sendResult(w http.ResponseWriter, httpstatus int, sequence []uint64, errormsg string) {

	results := fibSeqResult{httpstatus, sequence, errormsg}
	w.WriteHeader(results.HTTPStatus)
	json.NewEncoder(w).Encode(results)

}

// FibSeq - Calculate the fibonacci sequence out to 'index' numbers
func FibSeq(w http.ResponseWriter, r *http.Request) {

	var i int

	queryValues := r.URL.Query()

	// If more than one URI parameter has been passed, it is a bad request.
	if len(queryValues) != 1 {
		sendResult(w, http.StatusBadRequest, nil, fmt.Sprintf("Expected 1 parameter, got %d.", len(queryValues)))
		return
	}

	// If the URI parameter index has not been specified, it is a bad request.
	if queryValues.Get("index") == "" {
		sendResult(w, http.StatusBadRequest, nil, fmt.Sprintf("Invalid parameter. Expected 'index'"))
		return
	}

	// Is the value for the index parameter and integer, if not it is a bad request.
	val, err := strconv.Atoi(queryValues.Get("index"))
	if err != nil {
		sendResult(w, http.StatusBadRequest, nil, fmt.Sprintf("Invalid value sent for index : %v", queryValues.Get("index")))
		return
	}

	// Account for starting at 1, not 0
	val = val - 1

	if val >= -1 {
		f := make([]uint64, val+1)

		// Iterate through and calculate fibonaci sequence
		for i = 0; i <= val; i++ {

			switch i {
			// First two numbers can't be calculated, so set them.
			case 0, 1:
				f[i] = uint64(i)
			// Calculate the remaining numbers in the sequence
			default:
				f[i] = f[i-1] + f[i-2]
			}
		}
		sendResult(w, http.StatusOK, f, "")
	} else {
		sendResult(w, http.StatusBadRequest, nil, "Invalid index give, cannot be negative")
	}

}
