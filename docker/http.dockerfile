FROM golang:1.13.8-alpine3.11 AS build
MAINTAINER Adi Saripuloh <adisaripuloh@gmail.com>
# Install tools required to build the project
# We will need to run `docker build --no-cache .` to update those dependencies
RUN apk update \
	&& apk add --no-cache git
# go.mod and go.sum lists project dependencies
COPY go.mod go.sum /go/src/github.com/AdiSaripuloh/goproductapi/
WORKDIR /go/src/github.com/AdiSaripuloh/goproductapi/
# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download
# Remove tools after build the project
RUN apk del git
# Copy all project and build it
# This layer will be rebuilt when ever a file has changed in the project directory
COPY . /go/src/github.com/AdiSaripuloh/goproductapi/
RUN cd cmd/http/ \
	&& CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main

# This results in a single layer image
FROM alpine:latest
RUN apk update && apk --no-cache add ca-certificates
COPY --from=build /go/src/github.com/AdiSaripuloh/goproductapi .
RUN chmod +x cmd/http/main
CMD ["./cmd/http/main"]
