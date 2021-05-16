package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIsValidCheckSum(t *testing.T) {
	Convey("valid check sum", t, func() {
		testIban := &IBAN{CountryCode: "SE", CheckDigits: "45", BBAN: "50000000058398257466"}
		actualRes := testIban.isValidCheckSum()
		So(actualRes, ShouldBeTrue)
	})
	Convey("invalid check digits", t, func() {
		testIban := &IBAN{CountryCode: "SE", CheckDigits: "50", BBAN: "50000000058398257466"}
		actualRes := testIban.isValidCheckSum()
		So(actualRes, ShouldBeFalse)
	})
	Convey("invalid reminder", t, func() {
		testIban := &IBAN{CountryCode: "SE", CheckDigits: "45", BBAN: "50000000058398257468"}
		actualRes := testIban.isValidCheckSum()
		So(actualRes, ShouldBeFalse)
	})
}

func TestIsValidCountryFormat(t *testing.T) {
	Convey("valid country format", t, func() {
		testIban := &IBAN{CountryCode: "SE", CheckDigits: "45", BBAN: "50000000058398257466"}
		actualRes := testIban.isValidCountryFormat()
		So(actualRes, ShouldBeTrue)
	})
	Convey("invalid BBAN length", t, func() {
		testIban := &IBAN{CountryCode: "SE", CheckDigits: "45", BBAN: "5000000005839825"}
		actualRes := testIban.isValidCountryFormat()
		So(actualRes, ShouldBeFalse)
	})
	Convey("invalid country format", t, func() {
		testIban := &IBAN{CountryCode: "SE", CheckDigits: "45", BBAN: "5000000005839825TEST"}
		actualRes := testIban.isValidCountryFormat()
		So(actualRes, ShouldBeFalse)
	})
}

func TestIsValidFormat(t *testing.T) {
	Convey("invalid country code", t, func() {
		testIbanStr := "5558DOHB00001234567890ABCDEFG"
		actualRes := isValidFormat(testIbanStr, ibanCommonFormat)
		So(actualRes, ShouldBeFalse)
	})
	Convey("invalid check digits", t, func() {
		testIbanStr := "QAQADOHB00001234567890ABCDEFG"
		actualRes := isValidFormat(testIbanStr, ibanCommonFormat)
		So(actualRes, ShouldBeFalse)
	})
	Convey("too short", t, func() {
		testIbanStr := "QA58DOHB"
		actualRes := isValidFormat(testIbanStr, ibanCommonFormat)
		So(actualRes, ShouldBeFalse)
	})
	Convey("valid common format", t, func() {
		testIbanStr := "QA58DOHB00001234567890ABCDEFG"
		actualRes := isValidFormat(testIbanStr, ibanCommonFormat)
		So(actualRes, ShouldBeTrue)
	})
}

func TestReorderCharacters(t *testing.T) {
	Convey("correct string", t, func() {
		testIban := &IBAN{CountryCode: "SE", CheckDigits: "45", BBAN: "50000000058398257466"}
		actualRes := reorderCharacters(testIban)
		expectedRes := "50000000058398257466SE45"
		So(actualRes, ShouldEqual, expectedRes)
	})
}

func TestReplaceLettersToDigits(t *testing.T) {
	Convey("short string success", t, func() {
		testStr := "ABZ"
		actualRes := replaceLettersToDigits(testStr)
		expectedRes := "101135"
		So(actualRes, ShouldEqual, expectedRes)
	})
	Convey("iban string success", t, func() {
		testStr := "DOHB00001234567890ABCDEFGQA58"
		actualRes := replaceLettersToDigits(testStr)
		expectedRes := "132417110000123456789010111213141516261058"
		So(actualRes, ShouldEqual, expectedRes)
	})
}

func TestStringToIban(t *testing.T) {
	Convey("converted successfully", t, func() {
		testIbanStr := "QA58DOHB00001234567890ABCDEFG"
		actualRes := stringToIban(testIbanStr)
		expectedRes := &IBAN{CountryCode: "QA", CheckDigits: "58", BBAN: "DOHB00001234567890ABCDEFG"}
		So(actualRes, ShouldResemble, expectedRes)
	})
}
