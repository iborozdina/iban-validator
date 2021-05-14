package main

import (
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	response := &Response{
		StatusCode: http.StatusOK,
		Message:    "To use this service, send a GET request to '/validate' with the IBAN number to validate in 'iban' field value",
	}
	returnResponse(w, response)
}

func validateIBAN(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Println("Wrong request method sent:", r.Method)
		response := &Response{
			StatusCode: http.StatusMethodNotAllowed,
			Message:    "Method Not Allowed, should be a GET request",
		}
		returnResponse(w, response)
		return
	}

	// only one parameter is expected
	req := r.URL.Query().Get("iban")
	log.Println("GET param is:", req)

	// convert the requested string to IBAN struct
	iban, err := stringToIban(req)
	if err != nil {
		returnResponse(w, &Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
		return
	}

	// check if a specified country is supported by IBAN validator
	if _, ok := Countries[iban.CountryCode]; !ok {
		returnResponse(w, &Response{
			StatusCode: http.StatusNotFound,
			Message:    "Such country is not supported yet",
		})
		return
	}

	// check the format for specific country
	if !iban.isValidCountryStructure() {
		returnResponse(w, &Response{
			StatusCode: http.StatusBadRequest,
			Message:    "BBAN structure does not correspond the country format",
		})
		return
	}

	// validate the Check Digits
	// NOT IMPLEMENTED: validate the National Check Digits (used within the BBAN)
	if !iban.areValidCheckDigits() {
		returnResponse(w, &Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Check Digits are not valid",
		})
		return
	}

	response := &Response{
		StatusCode:       http.StatusOK,
		IBAN:             iban.CountryCode + iban.CheckDigits + iban.BBAN,
		CheckDigits:      iban.CheckDigits,
		Country:          Countries[iban.CountryCode].Name,
		CountryCode:      iban.CountryCode,
		ValidationResult: true,
	}
	returnResponse(w, response)
}
