FROM golang:1.13.8-alpine3.11 AS build
MAINTAINER Adi Saripuloh <adisaripuloh@gmail.com>
# set GOBIN path
# Install tools required to build the project
# We will need to run `docker build --no-cache .` to update those dependencies
RUN export GOBIN=$GOPATH/bin \
	&& apk add --no-cache git curl \
	&& curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
# Gopkg.toml and Gopkg.lock lists project dependencies
# These layers will only be re-built when Gopkg files are updated
COPY Gopkg.lock Gopkg.toml /go/src/github.com/AdiSaripuloh/goproductapi/
WORKDIR /go/src/github.com/AdiSaripuloh/goproductapi/
# Install library dependencies
RUN dep ensure -v -vendor-only
# Remove tools after build the project
RUN apk del git curl
# Copy all project and build it
# This layer will be rebuilt when ever a file has changed in the project directory
COPY . /go/src/github.com/AdiSaripuloh/goproductapi/
RUN cd cmd/http/ \
	&& go build -o main

# This results in a single layer image
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=build /go/src/github.com/AdiSaripuloh/goproductapi .
RUN chmod +x cmd/http/main
CMD ["./cmd/http/main"]
