package main

import (
	"log"
	"math/big"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var ibanCommonFormat = `^[A-Z]{2}\d{2}[A-Z0-9]{11,30}$`

type IBAN struct {
	CountryCode string
	CheckDigits string
	BBAN        string
}

// Get the full IBAN string
func (iban *IBAN) get() string {
	return iban.CountryCode + iban.CheckDigits + iban.BBAN
}

// Validate the Check Digits by reordering characters in IBAN, converting letters to digits and computing the reminder for mod 97
func (iban *IBAN) isValidCheckSum() bool {
	// reorder the characters of IBAN
	reorderedStr := reorderCharacters(iban)
	// replace letters with digits
	replacedLettersStr := replaceLettersToDigits(reorderedStr)
	// convert IBAN to big integer (*big.Int) for further calculations
	intIBAN := convertIbanStringToBigInt(replacedLettersStr)
	if intIBAN == nil {
		// can be changed to return internal error, it will be more relevant
		log.Printf("Could not convert IBAN string to big int")
		return false
	}

	// compute the remainder of IBAN(*big.Int) on division by 97 and check if reminder is equal to 1
	reminder := new(big.Int).Mod(intIBAN, big.NewInt(97))
	if reminder.Int64() != 1 {
		return false
	}
	return true
}

// Validate if BBAN of this IBAN corresponds to the specific country format and has the right length
func (iban *IBAN) isValidCountryFormat() bool {
	cc := iban.CountryCode
	if len(iban.BBAN) == Countries[cc].BBANLength && isValidFormat(iban.BBAN, Countries[cc].BBANFormat) {
		return true
	}
	return false
}

// Replace letters with digits and convert to big integer
func convertIbanStringToBigInt(iban string) *big.Int {
	result, ok := new(big.Int).SetString(iban, 10)
	if !ok {
		// can be changed to return internal error, it will be more relevant
		log.Printf("Could not convert IBAN to int")
		return nil
	}
	
	return result
}

// Validate the format of the string (str) according to provided regexp (requiredFormat)
func isValidFormat(str, requiredFormat string) bool {
	return regexp.MustCompile(requiredFormat).MatchString(str)
}

// Move the four initial characters to the end of the string
func reorderCharacters(iban *IBAN) string {
	return iban.BBAN + iban.CountryCode + iban.CheckDigits
}

// Replace each letter in the string with two digits, thereby expanding the string, where A = 10, B = 11, ..., Z = 35
func replaceLettersToDigits(str string) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var result string
	var num string

	for _, r := range str {
		if !unicode.IsDigit(r) {
			for i, v := range charset {
				if v == r {
					num = strconv.Itoa(i + 10)
					break
				}
			}
		} else {
			num = string(r)
		}
		result = result + num
	}
	return result
}

// Convert the string to IBAN structure: Country Code, Check Digits, BBAN
func stringToIban(str string) *IBAN {
	// remove spaces and force uppercase
	str = strings.ToUpper(strings.Replace(str, " ", "", -1))

	// validate common format and return IBAN struct
	if isValidFormat(str, ibanCommonFormat) {
		return &IBAN{
			CountryCode: str[0:2],
			CheckDigits: str[2:4],
			BBAN:        str[4:],
		}
	}
	return nil
}
