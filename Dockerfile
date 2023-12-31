FROM golang:1.21-alpine

ADD . /go/src

WORKDIR /go/src

RUN go get -d -v ./...
RUN go install -v ./...

RUN apk add build-base