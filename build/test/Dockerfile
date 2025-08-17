FROM golang:1.24.5-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

EXPOSE 8082
