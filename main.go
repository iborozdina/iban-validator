package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const port = ":8080"

func main() {
	log.Printf("Starting service...")
	r := mux.NewRouter()
	r.HandleFunc("/validate/{iban}", validateIBAN).Methods("GET")
	http.Handle("/", r)

	log.Printf("Serving on %s", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func returnResponse(w http.ResponseWriter, code int, r interface{}) {
	JSON, err := json.Marshal(r)
	if err != nil {
		log.Printf("Could not convert to JSON: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(JSON)
	if err != nil {
		log.Printf("Failed to write response body: %s", err)
	}
}
