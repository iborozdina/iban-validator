package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type (
	serverConfigStruct struct {
		host string
		port string
	}
	Response struct {
		StatusCode       int    `json:"statusCode"`
		IBAN             string `json:"iban,omitempty"`
		CheckDigits      string `json:"checkDigits,omitempty"`
		Country          string `json:"country,omitempty"`
		CountryCode      string `json:"countryCode,omitempty"`
		ValidationResult bool   `json:"valid"`
		Message          string `json:"message,omitempty"`
	}
)

var serverConfig = serverConfigStruct{
	host: "localhost",
	port: "8080",
}

func returnResponse(w http.ResponseWriter, r *Response) {
	JSON, err := json.Marshal(r)
	if err != nil {
		// todo: think if I need a log.Fatal here (exit)
		log.Println("Could not convert to JSON", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.StatusCode)
	w.Write(JSON)
}

func startServer() {
	log.Println("Start service")
	http.HandleFunc("/", index)
	http.HandleFunc("/validate", validateIBAN)

	log.Printf("%s\n", "Serving on http://"+serverConfig.host+":"+serverConfig.port+"/")
	err := http.ListenAndServe(serverConfig.host+":"+serverConfig.port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
