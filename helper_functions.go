package main

import (
	"log"
	"math/big"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var ibanFormat = `^[A-Z]{2}\d{2}[A-Z0-9]{11,30}$`

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
	// reorder the characters of IBAN and convert IBAN to big integer (*big.Int) to further calculations
	intIBAN := convertIbanStringToBigInt(reorderCharacters(iban))
	if intIBAN == nil {
		log.Println("Could not convert IBAN string to big int")
		return false
	}

	// compute the remainder of IBAN(*big.Int) on division by 97 and check if reminder is equal to 1
	reminder := new(big.Int).Mod(intIBAN, big.NewInt(97))
	log.Println("Reminder for IBAN mod 97 is: " + reminder.String())
	if reminder.Int64() == 1 {
		return true
	}

	return false
}

// Validate if BBAN of this IBAN corresponds to the specific country format and has the right length
func (iban *IBAN) isValidCountryFormat() bool {
	cc := iban.CountryCode
	if len(iban.BBAN) == Countries[cc].BBANLength && isValidFormat(iban.BBAN, Countries[cc].BBANFormat) {
		log.Println("IBAN corresponds the country format")
		return true
	}
	return false
}

// Replace each letter in the string with two digits, thereby expanding the string, where A = 10, B = 11, ..., Z = 35 and convert to big integer
func convertIbanStringToBigInt(iban string) *big.Int {
	// replace letters to digits
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var ibanAllNumbers string
	var num string

	for _, r := range iban {
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
		ibanAllNumbers = ibanAllNumbers + num
	}
	log.Println("Replaced all letters in IBAN: " + ibanAllNumbers)

	// convert string to big integer
	result, ok := new(big.Int).SetString(ibanAllNumbers, 10)
	if !ok {
		log.Println("Could not convert IBAN to int")
		return nil
	}
	log.Println("Converted IBAN to big integer: " + result.String())

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

// Convert the string to IBAN structure: Country Code, Check Digits, BBAN
func stringToIban(str string) *IBAN {
	// remove spaces and force uppercase
	str = strings.ToUpper(strings.Replace(str, " ", "", -1))

	// validate common format and return IBAN struct
	if isValidFormat(str, ibanFormat) {
		log.Println("IBAN corresponds the common format")
		return &IBAN{
			CountryCode: str[0:2],
			CheckDigits: str[2:4],
			BBAN:        str[4:],
		}
	}
	return nil
}
