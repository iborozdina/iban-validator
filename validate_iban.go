package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Response struct {
	Valid bool `json:"valid"`
}

type ErrResponse struct {
	Message string `json:"message"`
}

func validateIBAN(w http.ResponseWriter, r *http.Request) {
	ibanParam := mux.Vars(r)["iban"]
	log.Printf("Requested to validate IBAN with param: %s", ibanParam)

	if len(ibanParam) == 0 {
		returnResponse(w, http.StatusBadRequest, &ErrResponse{Message: "No IBAN number was provided in request"})
		return
	}

	// convert the requested string to IBAN struct, if the string corresponds to the common IBAN format
	iban := stringToIban(ibanParam)
	if iban == nil {
		returnResponse(w, http.StatusBadRequest, &ErrResponse{Message: "No IBAN number was provided in request"})
		return
	}

	// check if a specified country is supported by IBAN validator
	if _, ok := Countries[iban.CountryCode]; !ok {
		returnResponse(w, http.StatusNotFound, &ErrResponse{"The country is not supported yet"})
		return
	}

	// validate the format for specific country
	if !iban.isValidCountryFormat() {
		returnResponse(w, http.StatusOK, &Response{Valid: false})
		return
	}

	// validate the Check Digits
	if !iban.isValidCheckSum() {
		returnResponse(w, http.StatusOK, &Response{Valid: false})
		return
	}

	// todo: validate the National Check Digits (used within the BBAN)

	// return successful response for valid IBAN
	returnResponse(w, http.StatusOK, &Response{Valid: true})
}
