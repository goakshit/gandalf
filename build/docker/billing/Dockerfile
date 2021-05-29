FROM golang:1.14.3-alpine as builder

# Allows to filter & cleanup intermediate build images later
LABEL stage=builder
LABEL maintainer="Akshit Dhar"

ENV GO111MODULE=on

RUN mkdir /code

WORKDIR /code

COPY go.mod .
COPY go.sum .

RUN apk add build-base && \
    apk --update add ca-certificates && \
    go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -tags musl -o /code/exec ./cmd/billing
CMD ["/code/exec"]