package tests

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/doug-martin/goqu/v9"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var MainURL string

// db is for access the database in tests. it need for deleting rows created/updated in the relevants tests
var db *goqu.Database

func TestMain(m *testing.M) {
	testDbUrl := "postgres://test_user:test_password@test_database:5432/test_database?sslmode=disable"
	cli, err := sql.Open("postgres", testDbUrl)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	defer cli.Close()

	db = goqu.New("postgres", cli)
	mg, err := migrate.New("file://../migration_files", testDbUrl)
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
	setUpDb(db)
	slog.Info("Database setup\n\n")

	slog.Info("Start test server\n\n")
	server := httptest.NewServer(getTestServer(db))
	defer server.Close()

	MainURL = server.URL + "/test"
	slog.Info("Test server runned", "MainURL", MainURL)
	fmt.Print("\n")

	slog.Info("Start tests\n\n")
	code := m.Run()
	fmt.Print("\n")
	slog.Info("Tests finished\n\n")

	slog.Info("Start dropping migrations in test database")
	mg.Drop()
	slog.Info("Migrations dropped\n\n")

	os.Exit(code)
}
