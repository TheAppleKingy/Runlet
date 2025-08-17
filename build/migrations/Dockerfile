FROM golang:1.24.5-alpine

COPY ./ /migration_files/

WORKDIR /migration_files

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN apk add git build-base postgresql-dev
 

