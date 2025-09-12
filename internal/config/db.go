package config

import (
	"errors"
	"fmt"
	"os"
)

func GetDBUrl() (string, error) {
	dbName := os.Getenv("POSTGRES_DB")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	if dbName == "" {
		return "", errors.New("no db name was set in env")
	}
	if dbUser == "" {
		return "", errors.New("no db user was set in env")
	}
	if dbPassword == "" {
		return "", errors.New("no db password was set in env")
	}
	return fmt.Sprintf("postgres://%s:%s@database:5432/%s?sslmode=disable", dbUser, dbPassword, dbName), nil
}
