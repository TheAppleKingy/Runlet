FROM golang:1.24.5-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY . .

RUN swag init -g ./cmd/server/main.go    
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/server

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/app .


EXPOSE 8081
