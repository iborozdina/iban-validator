FROM golang:1.16.4-alpine3.13

WORKDIR /go/src/app
COPY . .

RUN go build -o bin/main .

CMD ["bin/main"]
