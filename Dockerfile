# syntax=docker/dockerfile:1
FROM golang:1.16-alpine

RUN apk add --no-cache make

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN go build
