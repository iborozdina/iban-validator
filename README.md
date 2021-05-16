###### **Just a test case, but working**

_____


# iban-validator
A simple web-service in Golang for validating IBAN numbers.


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
> `curl localhost:8080/validate/  QA 58 DOHB 00001234567890 ABCDEF G`

  
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

  
  
## About International Bank Account Number
To read more about IBAN, please visit [this Wikipedia page](https://en.wikipedia.org/wiki/International_Bank_Account_Number)
