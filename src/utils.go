package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJson(w http.ResponseWriter, statusCode int, payload interface{}) {

	dat, err := json.Marshal(payload)

	if err != nil {
		w.WriteHeader(500)
		log.Printf("failed to marshal payload to json: %v", payload)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(dat)
}

func respondWithError(w http.ResponseWriter, statusCode int, msg string) {

	if statusCode > 499 {
		log.Printf("responding with 5XX code: %+v", msg)
	}

	type ErrorResp struct {
		Error string `json:"error"`
	}

	respondWithJson(w, statusCode, ErrorResp{Error: msg})
}

func readJson(r *http.Request, payload interface{}) error {

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(payload)

	if err != nil {
		log.Printf("failed to decode json : %+v", err)
		return err
	}

	return nil
}