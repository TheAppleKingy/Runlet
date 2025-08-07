package main

import (
	"Runlet/internal/infrastructure"
	"context"
	"log/slog"
	"os"
	"time"
)

func main() {
	dbClient := infrastructure.GetDbClient()
	defer dbClient.Close()
	startConnectingTime := time.Now()
	for {
		if time.Since(startConnectingTime) > 6*time.Second {
			slog.Error("Migrations was not applied")
			os.Exit(1)
		}
		err := dbClient.Schema.Create(context.Background())
		if err == nil {
			slog.Info("Migrations applied successfully")
			return
		}
		slog.Error("Error applying migrations", "error", err, "status", "retry")
		time.Sleep(1 * time.Second)
	}
}
