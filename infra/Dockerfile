FROM golang:1.18-alpine

WORKDIR /go/src/app

RUN apk update && \
    apk add git && \
    go install github.com/togatoga/procon-gardener@latest