package main

import (
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	log.Println("Index endpoint is called successfully")
	response := &Response{
		StatusCode: http.StatusOK,
		Message:    "To use this service, send a GET request to '/validate' with the IBAN number to validate in 'iban' field value",
	}
	returnResponse(w, response)
}

func validateIBAN(w http.ResponseWriter, r *http.Request) {
	// check the request method
	if r.Method != "GET" {
		log.Println("Wrong request method sent:", r.Method)
		response := &Response{
			StatusCode: http.StatusMethodNotAllowed,
			Message:    "Method Not Allowed, should be a GET request",
		}
		returnResponse(w, response)
		return
	}
	log.Println("Validate endpoint is called...")

	// only one parameter is expected in request
	req := r.URL.Query().Get("iban")
	log.Println("GET parameter is:", req)
	if len(req) == 0 {
		message := "No IBAN number was provided in request"
		log.Println(message)
		returnResponse(w, &Response{
			StatusCode: http.StatusBadRequest,
			Message:    message,
		})
		return
	}

	// convert the requested string to IBAN struct
	iban := stringToIban(req)
	if iban == nil {
		message := "Could not convert the IBAN from string to IBAN struct, does not correspond to the common IBAN format. Request parameter: " + req
		log.Println(message)
		returnResponse(w, &Response{
			StatusCode:       http.StatusBadRequest,
			ValidationResult: false,
			Message:          message,
		})
		return
	}

	// check if a specified country is supported by IBAN validator
	if _, ok := Countries[iban.CountryCode]; !ok {
		message := "Such country is not supported yet. Country code: " + iban.CountryCode
		log.Println(message)
		returnResponse(w, &Response{
			StatusCode:       http.StatusNotFound,
			IBAN:             iban.get(),
			CountryCode:      iban.CountryCode,
			ValidationResult: false,
			Message:          message,
		})
		return
	}

	// check the format for specific country
	if !iban.isValidCountryFormat() {
		message := "BBAN structure does not correspond the country format. IBAN: " + iban.get()
		log.Println(message)
		returnResponse(w, &Response{
			StatusCode:       http.StatusBadRequest,
			IBAN:             iban.get(),
			CheckDigits:      iban.CheckDigits,
			CountryCode:      iban.CountryCode,
			Country:          Countries[iban.CountryCode].Name,
			ValidationResult: false,
			Message:          message,
		})
		return
	}

	// validate the Check Digits
	// NOT IMPLEMENTED: validate the National Check Digits (used within the BBAN)
	if !iban.isValidCheckSum() {
		message := "CheckSum is not valid. IBAN: " + iban.get()
		log.Println(message)
		returnResponse(w, &Response{
			StatusCode:       http.StatusBadRequest,
			IBAN:             iban.get(),
			CheckDigits:      iban.CheckDigits,
			CountryCode:      iban.CountryCode,
			Country:          Countries[iban.CountryCode].Name,
			ValidationResult: false,
			Message:          message,
		})
		return
	}

	// return successful response for valid IBAN
	log.Println("IBAN is valid!")
	response := &Response{
		StatusCode:       http.StatusOK,
		IBAN:             iban.get(),
		CheckDigits:      iban.CheckDigits,
		CountryCode:      iban.CountryCode,
		Country:          Countries[iban.CountryCode].Name,
		ValidationResult: true,
	}
	returnResponse(w, response)
}
