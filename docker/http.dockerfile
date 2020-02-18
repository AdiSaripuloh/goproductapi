FROM golang:1.13.8-alpine3.11

LABEL maintainer="Adi Saripuloh <adisaripuloh@gmail.com>"

COPY . /go/src/github.com/AdiSaripuloh/goproductapi/

VOLUME .:/go/src/github.com/AdiSaripuloh/goproductapi/

RUN export GOBIN=$GOPATH/bin

WORKDIR /go/src/github.com/AdiSaripuloh/goproductapi/

RUN apk update \
	&& apk add --no-cache curl git mercurial \
	&& curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh \
	&& dep ensure -v --vendor-only \
	&& apk del git mercurial

RUN cd cmd/http/ \
	&& go build -o main \
	&& cd ../../

CMD ["./cmd/http/main"]
