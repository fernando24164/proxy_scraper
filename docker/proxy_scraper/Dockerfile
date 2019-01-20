FROM golang:1.11.1-alpine3.8

COPY . /go/src/proxy_scraper

WORKDIR /go/src/proxy_scraper

RUN apk add --no-cache curl git \
    && curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh \
    && dep ensure \
    && go build;

ENTRYPOINT go run main.go
