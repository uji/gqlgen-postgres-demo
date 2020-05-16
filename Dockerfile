FROM golang:1.14.3-alpine3.11
RUN mkdir /go/src/repo
WORKDIR /go/src/repo
ADD . /go/src/repo
