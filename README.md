###### **Just a test case, but working**

_____

# iban-validator

A simple web-service in Golang for validating IBAN numbers.

### About International Bank Account Number

The International Bank Account Number (IBAN) is an internationally agreed system of identifying bank accounts across
national borders to facilitate the communication and processing of cross border transactions with a reduced risk of
transcription errors.

The IBAN consists of up to 34 alphanumeric characters, as follows:

* country code using ISO 3166-1 alpha-2 – two letters,
* check digits – two digits,
* Basic Bank Account Number (BBAN) – up to 30 alphanumeric characters that are country-specific.

To read more about IBAN, please
visit [this Wikipedia page](https://en.wikipedia.org/wiki/International_Bank_Account_Number)

### Validating the IBAN

Check that the provided parameter has the correct common IBAN structure (country code - 2 letters, check digits - 2
digits, BBAN - 11-30 alpha-numeric characters, white-spaces are ignored). If not, you get the message "No IBAN number
was provided in request".

An IBAN is validated by converting it into an integer and performing a basic mod-97 operation on it. If the IBAN is
valid, the remainder equals 1. The algorithm of IBAN validation is as follows:

* Check that the total IBAN length is correct as per the country. If not, the IBAN is invalid.
* Move the four initial characters to the end of the string.
* Replace each letter in the string with two digits, thereby expanding the string, where A = 10, B = 11, ..., Z = 35.
* Interpret the string as a decimal integer and compute the remainder of that number on division by 97.
* If the remainder is 1, the check digit test is passed and the IBAN **might be valid**.

### National Check Digits

In addition to the IBAN check digits, some countries have their own national check digits used within the BBAN, as part
of their national account number formats. Each country determines its own algorithm used for assigning and validating
the national check digits - some relying on international standards, some inventing their own national standard, and
some allowing each bank to decide if or how to implement them. Some algorithms apply to the entire BBAN, and others to
one or more of the fields within it. The check digits may be considered an integral part of the account number, or an
external field separate from the account number, depending on the country's rules.

**PLEASE NOTE!** This National Check Digits validation is NOT implemented in this validator yet.

## Running the app

Clone the repo and change the directory to `iban-validator`

Run with make:

```
make build && make run
```

Run with docker:

```
docker-compose up
```

To re-build and run the app with docker:

```
docker-compose up --build
```

The service will launch and listen on the **port 8080**.

#### Running the tests

```
make test
```

## Usage

To validate your IBAN, call the `validate` endpoint with parameter from your browser:

```
http://localhost:8080/validate/SE4550000000058398257466
```

Or using `curl`:

```
curl localhost:8080/validate/SE4550000000058398257466
```

**PLEASE NOTE!**
> Using `curl` you should encode your IBAN to URL-encoded format: <br>
> GOOD:
> `curl localhost:8080/validate/%20%20QA%2058%20DOHB%2000001234567890%20ABCDEF%20G` <br>
> NOT GOOD:
> `curl localhost:8080/validate/ QA 58 DOHB 00001234567890 ABCDEF G`

## What to expect in Response

In response you will get a very simple JSON.

* If IBAN is valid:

```
Status Code: 200 OK
Body: {"valid":true}
```

* If IBAN is NOT valid:

```
Status Code: 200 OK
Body: {"valid":false}
```

* If nothing was provided as parameter or the parameter does not look like IBAN:

```
Status Code: 400 Bad Request
Body: {"message":"No IBAN number was provided in request"}
```

* If the country of provided IBAN is not supported by this validator:

```
Status Code: 404 Not Found
Body: {"message":"The country is not supported yet"}
```


