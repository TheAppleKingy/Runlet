package infrastructure

import (
	"Runlet/internal/infrastructure/ent"
	"log/slog"
	"os"

	_ "github.com/lib/pq"
)

func GetDbClient() *ent.Client {
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		slog.Error("Database url did not set in environment")
		os.Exit(1)
	}
	client, err := ent.Open("postgres", dbUrl)
	if err != nil {
		slog.Error("error database connection", "error", err)
		os.Exit(1)
	}
	return client
}
