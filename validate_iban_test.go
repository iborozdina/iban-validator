package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	. "github.com/smartystreets/goconvey/convey"
)

func TestValidateIBAN(t *testing.T) {
	Convey("no param provided", t, func() {
		r, _ := http.NewRequest("GET", "/validate", nil)
		w := httptest.NewRecorder()

		validateIBAN(w, r)

		var m map[string]interface{}
		_ = json.Unmarshal(w.Body.Bytes(), &m)

		So(w.Code, ShouldEqual, http.StatusBadRequest)
		So(m["message"], ShouldEqual, "No IBAN number was provided in request")
	})
	Convey("invalid param with spaces", t, func() {
		r, _ := http.NewRequest("GET", "/validate", nil)
		w := httptest.NewRecorder()
		r = mux.SetURLVars(r, map[string]string{"iban": "  "})

		validateIBAN(w, r)

		var m map[string]interface{}
		_ = json.Unmarshal(w.Body.Bytes(), &m)

		So(w.Code, ShouldEqual, http.StatusBadRequest)
		So(m["message"], ShouldEqual, "No IBAN number was provided in request")
	})
	Convey("invalid param", t, func() {
		r, _ := http.NewRequest("GET", "/validate", nil)
		w := httptest.NewRecorder()
		r = mux.SetURLVars(r, map[string]string{"iban": "123 456"})

		validateIBAN(w, r)

		var m map[string]interface{}
		_ = json.Unmarshal(w.Body.Bytes(), &m)

		So(w.Code, ShouldEqual, http.StatusBadRequest)
		So(m["message"], ShouldEqual, "No IBAN number was provided in request")
	})
	Convey("country not supported", t, func() {
		r, _ := http.NewRequest("GET", "/validate", nil)
		w := httptest.NewRecorder()
		r = mux.SetURLVars(r, map[string]string{"iban": "SS 12 1234567890123"})

		validateIBAN(w, r)

		var m map[string]interface{}
		_ = json.Unmarshal(w.Body.Bytes(), &m)

		So(w.Code, ShouldEqual, http.StatusNotFound)
		So(m["message"], ShouldEqual, "The country is not supported yet")
	})
	Convey("incorrect iban Norway", t, func() {
		r, _ := http.NewRequest("GET", "/validate", nil)
		w := httptest.NewRecorder()
		r = mux.SetURLVars(r, map[string]string{"iban": "NO 93 1234 567890 1"})

		validateIBAN(w, r)

		var m map[string]interface{}
		_ = json.Unmarshal(w.Body.Bytes(), &m)

		So(w.Code, ShouldEqual, http.StatusOK)
		So(m["valid"], ShouldBeFalse)
	})
	Convey("incorrect iban Sweden", t, func() {
		r, _ := http.NewRequest("GET", "/validate", nil)
		w := httptest.NewRecorder()
		r = mux.SetURLVars(r, map[string]string{"iban": "SE 00 50000000058398257466"})

		validateIBAN(w, r)

		var m map[string]interface{}
		_ = json.Unmarshal(w.Body.Bytes(), &m)

		So(w.Code, ShouldEqual, http.StatusOK)
		So(m["valid"], ShouldBeFalse)
	})
	Convey("incorrect iban Qatar", t, func() {
		r, _ := http.NewRequest("GET", "/validate", nil)
		w := httptest.NewRecorder()
		r = mux.SetURLVars(r, map[string]string{"iban": "  QA 58 DOHB 00001234567890 ABCDEF 0 "})

		validateIBAN(w, r)

		var m map[string]interface{}
		_ = json.Unmarshal(w.Body.Bytes(), &m)

		So(w.Code, ShouldEqual, http.StatusOK)
		So(m["valid"], ShouldBeFalse)
	})
	Convey("correct iban Sweden no spaces", t, func() {
		r, _ := http.NewRequest("GET", "/validate", nil)
		w := httptest.NewRecorder()
		r = mux.SetURLVars(r, map[string]string{"iban": "SE4550000000058398257466"})

		validateIBAN(w, r)

		var m map[string]interface{}
		_ = json.Unmarshal(w.Body.Bytes(), &m)

		So(w.Code, ShouldEqual, http.StatusOK)
		So(m["valid"], ShouldBeTrue)
	})
	Convey("correct iban Qatar with spaces", t, func() {
		r, _ := http.NewRequest("GET", "/validate", nil)
		w := httptest.NewRecorder()
		r = mux.SetURLVars(r, map[string]string{"iban": "  QA 58 DOHB 00001234567890 ABCDEF G "})

		validateIBAN(w, r)

		var m map[string]interface{}
		_ = json.Unmarshal(w.Body.Bytes(), &m)

		So(w.Code, ShouldEqual, http.StatusOK)
		So(m["valid"], ShouldBeTrue)
	})
}
