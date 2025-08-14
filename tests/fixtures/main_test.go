package fixtures

import (
	"database/sql"
	"log/slog"
	"os"
	"testing"

	"github.com/doug-martin/goqu/v9"
	"github.com/golang-migrate/migrate/v4"
)

func TestMain(m *testing.M) {
	testDbUrl := os.Getenv("TEST_DATABASE_URL")
	if testDbUrl == "" {
		slog.Error("cannot get conn str for test db")
		os.Exit(1)
	}
	cli, err := sql.Open("postgres", testDbUrl)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	defer cli.Close()
	db := goqu.New("postgres", cli)
	mg, err := migrate.New("file://./migrations/queries", testDbUrl)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	mg.Drop()
	if err := mg.Up(); err != nil && err != migrate.ErrNoChange {
		slog.Error(err.Error())
		os.Exit(1)
	}
	setUpDb(db)
	code := m.Run()
	mg.Drop()
	os.Exit(code)
}
