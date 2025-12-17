package integration

import (
	"Runlet/internal/infrastructure/config"
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"testing"

	"github.com/doug-martin/goqu/v9"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var MainURL string

func TestMain(m *testing.M) {
	if err := config.LoadConfigs(); err != nil {
		slog.Error("cannot load config data", "error", err)
		os.Exit(1)
	}
	testDbUrl := config.DBConfig.TestURL
	cli, err := sql.Open("postgres", testDbUrl)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	DB = goqu.New("postgres", cli)
	mg, err := migrate.New("file://../../migration_files", testDbUrl)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	slog.Info("Start apply migrations to test database")
	if err := mg.Up(); err != nil && err != migrate.ErrNoChange {
		slog.Error(err.Error())
		os.Exit(1)
	}
	slog.Info("Migrations applied\n\n")

	slog.Info("Start setup test database")
	setUpDb(DB)
	slog.Info("Database setup\n\n")

	slog.Info("Start test server\n\n")
	server := getTestHTTPServer(DB)
	defer server.Close()
	MainURL = server.URL + "/test"

	slog.Info("Test server runned", "MainURL", MainURL)
	fmt.Print("\n")

	slog.Info("Starting integration tests\n\n")
	code := m.Run()
	fmt.Print("\n")
	slog.Info("Integration tests finished\n\n")

	os.Exit(code)
}
