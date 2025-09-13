FROM golang:1.24.5-alpine AS builder

WORKDIR /app

COPY . .

ENV RUNNERS_CONF_PATH=/app/runners.yaml

RUN go mod download

EXPOSE 8082
