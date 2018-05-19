package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type FibSeqResult struct {
	HTTPStatus int    `json:"httpstatus"`
	Sequence   []int  `json:"sequence"`
	ErrorMsg   string `json:"errormsg"`
}

// FibSeq
func FibSeq(w http.ResponseWriter, r *http.Request) {

	results := FibSeqResult{}

	queryValues := r.URL.Query()

	if len(queryValues) != 1 {
		results.HTTPStatus = http.StatusBadRequest
		errorMsg := fmt.Sprintf("Expected 1 parameter, got %d.", len(queryValues))
		results.ErrorMsg = errorMsg
	} else {

		if queryValues.Get("index") == "" {
			results.HTTPStatus = http.StatusBadRequest
			results.ErrorMsg = fmt.Sprintf("Invalid parameter. Expected 'index'")

		} else {

			val, _ := strconv.Atoi(queryValues.Get("index"))
			f := make([]int, val+1)
			var i int

			f[0] = 0
			f[1] = 1

			for i = 2; i <= val; i++ {
				f[i] = f[i-1] + f[i-2]
			}

			results.HTTPStatus = http.StatusOK
			results.Sequence = f
			results.ErrorMsg = ""
		}
	}

	w.WriteHeader(results.HTTPStatus)
	json.NewEncoder(w).Encode(results)

}
