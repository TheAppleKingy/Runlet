FROM golang:1.24.5-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY . .

RUN swag init -g ./cmd/server/main.go    
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/server

FROM scratch

WORKDIR /app

COPY --from=builder /app/app .
COPY --from=builder /app/runners.yaml .

ENV RUNNERS_CONF_PATH=/app/runners.yaml

EXPOSE 8081

ENTRYPOINT [ "/app/app" ]
