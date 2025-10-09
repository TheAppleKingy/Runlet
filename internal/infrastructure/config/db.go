package config

import (
	"fmt"
	"os"
)

type dbConfig struct {
	URL     string
	TestURL string
}

func (cfg *dbConfig) parse() error {
	dbName := os.Getenv("POSTGRES_DB")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	if dbName == "" {
		return ErrNoEnvDbName
	}
	if dbUser == "" {
		return ErrNoEnvUsername
	}
	if dbPassword == "" {
		return ErrNoEnvDbPassword
	}
	cfg.URL = fmt.Sprintf("postgres://%s:%s@database:5432/%s?sslmode=disable", dbUser, dbPassword, dbName)
	cfg.TestURL = "postgres://test_user:test_password@test_database:5432/test_database?sslmode=disable"
	return nil
}
