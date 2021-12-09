# syntax=docker/dockerfile:1
FROM golang:1.17-alpine

RUN apk add --no-cache \
        python3 \
        make \
        jq \
        py3-pip \
    && pip3 install --upgrade pip \
    && pip3 install --no-cache-dir \
        awscli \
    && rm -rf /var/cache/apk/*

RUN go get -u github.com/golang-migrate/migrate/v4@v4.15.1

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN go build
