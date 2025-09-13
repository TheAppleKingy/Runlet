package http

import (
	"fmt"
	"log/slog"
	"os"
	"testing"

	"Runlet/tests/integration"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var MainURL string

func TestMain(m *testing.M) {
	slog.Info("Start test server\n\n")
	server := GetTestHTTPServer(integration.DB)
	defer server.Close()
	MainURL = server.URL + "/test"

	slog.Info("Test server runned", "MainURL", MainURL)
	fmt.Print("\n")

	slog.Info("Start integration tests\n\n")
	code := m.Run()
	fmt.Print("\n")
	slog.Info("Integration tests finished\n\n")

	os.Exit(code)
}
