# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:alpine AS build

RUN apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o /app/pos main.go

##
## Deploy
##
FROM alpine

WORKDIR /app


COPY --from=build /app/pos /app/pos


ENTRYPOINT /app/pos