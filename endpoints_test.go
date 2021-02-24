package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPeople(t *testing.T){
	req,err := http.NewRequest("GET","/",nil)
	if err != nil {
		t.Fatal(err)
	}
	rr :=httptest.NewRecorder()
	handler := http.HandlerFunc(Index)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status !=http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	//Check the response body is what's expected
	expected := `[{"ID":1,"firstName":"Bianca","lastName":"Reusch","Age":27}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetPersonByID(t *testing.T){
	req, err := http.NewRequest("GET","/showPerson",nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("id","1")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(showPerson)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	//Check the response body is what's expected
	expected := `{"id":1,"FirstName":"Bianca","LastName":"Reusch","Age":27}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}