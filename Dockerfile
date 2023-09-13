FROM golang:1.21-alpine

ADD . /go/src

WORKDIR /go/src

RUN apk add build-base