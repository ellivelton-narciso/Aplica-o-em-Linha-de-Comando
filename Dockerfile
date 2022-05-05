FROM golang:latest

WORKDIR /go/src

COPY . /go/src

RUN apt update && apt install vim -y