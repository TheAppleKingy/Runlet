package service

import (
	"fmt"
	"log/slog"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Print("\n")
	slog.Info("Starting unit tests\n\n")
	m.Run()
	fmt.Print("\n")
	slog.Info("Unit tests finished\n\n")
}
