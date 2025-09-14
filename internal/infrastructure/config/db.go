package config

import (
	"fmt"
	"log/slog"
	"os"
)

type dbConfig struct {
	URL string
}

func (cfg *dbConfig) parse() {
	dbName := os.Getenv("POSTGRES_DB")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	if dbName == "" {
		slog.Error("no db name was set in env")
		os.Exit(1)
	}
	if dbUser == "" {
		slog.Error("no db user was set in env")
		os.Exit(1)
	}
	if dbPassword == "" {
		slog.Error("no db password was set in env")
		os.Exit(1)
	}
	cfg.URL = fmt.Sprintf("postgres://%s:%s@database:5432/%s?sslmode=disable", dbUser, dbPassword, dbName)
}

var DBConfig = dbConfig{}
