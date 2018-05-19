package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFibSeqHandler_NoParameters(t *testing.T) {

	req, err := http.NewRequest("GET", "/fibseq", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(FibSeq)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Status handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	expected := `{"httpstatus":400,"sequence":null,"errormsg":"Expected 1 parameter, got 0."}` + "\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got >%v< want >%v<",
			rr.Body.String(), expected)
	}

}

func TestFibSeqHandler_MultipleParamaters(t *testing.T) {

	req, err := http.NewRequest("GET", "/fibseq?index=5&foo=bar", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(FibSeq)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Status handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	expected := `{"httpstatus":400,"sequence":null,"errormsg":"Expected 1 parameter, got 2."}` + "\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got >%v< want >%v<",
			rr.Body.String(), expected)
	}

}

func TestFibSeqHandler_InvalidParamater(t *testing.T) {

	req, err := http.NewRequest("GET", "/fibseq?foo=bar", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(FibSeq)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Status handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	expected := `{"httpstatus":400,"sequence":null,"errormsg":"Invalid parameter. Expected 'index'"}` + "\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got >%v< want >%v<",
			rr.Body.String(), expected)
	}

}

func TestFibSeqHandler_ValidParamater(t *testing.T) {

	req, err := http.NewRequest("GET", "/fibseq?index=5", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(FibSeq)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"httpstatus":200,"sequence":[0,1,1,2,3,5],"errormsg":""}` + "\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got >%v< want >%v<",
			rr.Body.String(), expected)
	}
}
