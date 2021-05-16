build:
	go build -o bin/main .

test:
	go test

run:
	./bin/main

docker:
	docker build -t iban-validator:latest .
