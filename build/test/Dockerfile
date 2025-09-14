FROM golang:1.24.5-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV RUNNERS_CONF_PATH=/app/runners.yaml

EXPOSE 8082
