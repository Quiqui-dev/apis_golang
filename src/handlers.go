package main

import (
	"net/http"
)

type requestPayload struct {
	Num1 float64 `json:"num1"`
	Num2 float64 `json:"num2"`
}

type resultPayload struct {
	Result float64 `json:"result"`
}

func handleAdd(w http.ResponseWriter, r *http.Request) {
	
	var addParams = &requestPayload{}

	err := readJson(r, addParams)

	if err != nil {
		respondWithError(w, 400, "Could not decode json, bad request")
	}
	
	result := add(addParams.Num1, addParams.Num2)

	outPayload := resultPayload{
		Result: result,
	}

	respondWithJson(w, 200, outPayload)
}


func handleSubtract(w http.ResponseWriter, r *http.Request) {

	var subParams = &requestPayload{}

	err := readJson(r, subParams)

	if err != nil {
		respondWithError(w, 400, "Could not decode json, bad request")
	}

	result := subtract(subParams.Num1, subParams.Num2)

	outPayload := resultPayload{
		Result: result,
	}

	respondWithJson(w, 200, outPayload)
}


func handleMutliply(w http.ResponseWriter, r *http.Request) {

	var multParams = &requestPayload{}

	err := readJson(r, multParams)

	if err != nil {
		respondWithError(w, 400, "Could not decode json, bad request")
	}

	result := multiply(multParams.Num1, multParams.Num2)

	outPayload := resultPayload{
		Result: result,
	}

	respondWithJson(w, 200, outPayload)
}

func handleDivide(w http.ResponseWriter, r *http.Request) {

	var divideParams = &requestPayload{}

	err := readJson(r, divideParams)

	if err != nil {
		respondWithError(w, 400, "Could not decode json")
	}

	result, err := divide(divideParams.Num1, divideParams.Num2)

	if err != nil {
		respondWithError(w, 400, "Cannot divide by 0")
	}
	
	outPayload := resultPayload{
		Result: result,
	}

	respondWithJson(w, 200, outPayload)
}
