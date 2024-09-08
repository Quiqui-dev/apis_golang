package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestAddAPI(t *testing.T) {
	
	reqPayload := requestPayload{
		Num1: 5,
		Num2: 6,
	}

	jsonPayload, _ := json.Marshal(reqPayload)

	jsonBuf := bytes.NewBuffer(jsonPayload)

	req, err := http.NewRequest("POST", "/add", jsonBuf)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handleAdd(rr, req)

	assert.Equal(t, rr.Result().StatusCode, http.StatusOK)
}


func TestSubAPI(t *testing.T) {
	
	reqPayload := requestPayload{
		Num1: 5,
		Num2: 6,
	}

	jsonPayload, _ := json.Marshal(reqPayload)

	jsonBuf := bytes.NewBuffer(jsonPayload)

	req, err := http.NewRequest("POST", "/subtract", jsonBuf)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handleSubtract(rr, req)

	assert.Equal(t, rr.Result().StatusCode, http.StatusOK)
}


func TestMultAPI(t *testing.T) {
	
	reqPayload := requestPayload{
		Num1: 5,
		Num2: 6,
	}

	jsonPayload, _ := json.Marshal(reqPayload)

	jsonBuf := bytes.NewBuffer(jsonPayload)

	req, err := http.NewRequest("POST", "/multiply", jsonBuf)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handleMutliply(rr, req)

	assert.Equal(t, rr.Result().StatusCode, http.StatusOK)
}


func TestDivideAPI(t *testing.T) {
	
	reqPayload := requestPayload{
		Num1: 5,
		Num2: 6,
	}

	jsonPayload, _ := json.Marshal(reqPayload)

	jsonBuf := bytes.NewBuffer(jsonPayload)

	req, err := http.NewRequest("POST", "/divide", jsonBuf)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handleDivide(rr, req)

	assert.Equal(t, rr.Result().StatusCode, http.StatusOK)

	rr = httptest.NewRecorder()

	reqPayload.Num2 = 0

	jsonPayload, _ = json.Marshal(reqPayload)

	jsonBuf = bytes.NewBuffer(jsonPayload)

	req, err = http.NewRequest("POST", "/divide", jsonBuf)

	if err != nil {
		t.Fatal(err)
	}

	handleDivide(rr, req)

	assert.Equal(t, rr.Result().StatusCode, http.StatusBadRequest)

}