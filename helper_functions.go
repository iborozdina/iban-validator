package main

import (
	"log"
	"math/big"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/pkg/errors"
)

type IBAN struct {
	CountryCode string
	CheckDigits string
	BBAN        string
}

var ibanFormat = `^[A-Z]{2}\d{2}[A-Z0-9]{11,30}$`

// Converts the string to IBAN structure: Country Code string, Check Digits int, BBAN string
func stringToIban(str string) (*IBAN, error) {
	if len(str) == 0 {
		return nil, errors.New("no IBAN number was provided")
	}

	// Remove spaces and force uppercase
	str = strings.ToUpper(strings.Replace(str, " ", "", -1))

	if isValidFormat(str, ibanFormat) {
		//cg, err := strconv.Atoi(str[2:4])
		//if err != nil {
		//	return nil, err
		//}
		return &IBAN{
			CountryCode: str[0:2],
			CheckDigits: str[2:4],
			BBAN:        str[4:],
		}, nil
	}
	return nil, errors.New("IBAN structure is not valid")
}

// Validate the format of the string (str) according to provided regexp (requiredFormat)
func isValidFormat(str, requiredFormat string) bool {
	return regexp.MustCompile(requiredFormat).MatchString(str)
}

// Validate if BBAN of this IBAN corresponds to the specific country format and has the right length
func (iban *IBAN) isValidCountryStructure() bool {
	cc := iban.CountryCode
	if len(iban.BBAN) == Countries[cc].BBANLength && isValidFormat(iban.BBAN, Countries[cc].BBANFormat) {
		return true
	}
	return false
}

// Validate the Check Digits by reordering characters in IBAN, converting letters to digits and computing the reminder for mod 97
func (iban *IBAN) areValidCheckDigits() bool {

	// reorder the characters of IBAN and convert IBAN to big integer (*big.Int) to further calculations
	intIBAN := convertIbanStringToBigInt(reorderCharacters(iban))
	if intIBAN == nil {
		return false
	}

	// compute the remainder of IBAN(*big.Int) on division by 97 and check if reminder is equal to 1
	reminder := new(big.Int).Mod(intIBAN, big.NewInt(97))
	if reminder.Int64() == 1 {
		return true
	}

	return false
}

// Move the four initial characters to the end of the string
func reorderCharacters(iban *IBAN) string {
	return iban.BBAN + iban.CountryCode + iban.CheckDigits
}

// Replace each letter in the string with two digits, thereby expanding the string, where A = 10, B = 11, ..., Z = 35 and convert to *big.Int
func convertIbanStringToBigInt(iban string) *big.Int {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	var ibanAllNumbers string
	var num string
	var result *big.Int

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

	result, ok := new(big.Int).SetString(ibanAllNumbers, 10)
	if !ok {
		log.Println("Could not convert IBAN to int")
		return nil
	}

	return result
}
