FROM golang:1.24.5-alpine

COPY ../queries/ /migrations/

WORKDIR /migrations

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN apk add git build-base postgresql-dev
 

