ARG GO_VERSION=1.12.5-alpine3.9

FROM golang:$GO_VERSION as builder
MAINTAINER shane shane.h.1@gmail.com

COPY . assemblyPayments$GOPATH/src/assemblyPayments
WORKDIR $GOPATH/src/assemblyPayments
ADD . $GOPATH/src/assemblyPayments
COPY Gopkg.toml Gopkg.lock ./
# Install Git
RUN apk update && \
	apk add git
# Install dep
ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep
# Fetch dependencies.
RUN dep ensure --vendor-only
# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -i -v -o /go/bin/client assemblyPayments/cmd

