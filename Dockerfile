# syntax=docker/dockerfile:1
FROM golang:1.17-alpine

RUN apk add --no-cache \
        python3 \
        curl \
        vim \
        make \
        jq \
        bash \
        groff \
        py3-pip \
    && pip3 install --upgrade pip \
    && pip3 install --no-cache-dir \
        awscli \
    && rm -rf /var/cache/apk/*

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.1/migrate.linux-amd64.tar.gz \
    --output migrate.tar.gz \
    && tar xvf migrate.tar.gz migrate \
    && mv migrate /usr/local/bin/

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

COPY Makefile Makefile

COPY migrations/ ./migrations/

RUN go build
