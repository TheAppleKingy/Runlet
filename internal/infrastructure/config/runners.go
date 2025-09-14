package config

import (
	"log/slog"
	"os"

	"gopkg.in/yaml.v3"
)

type runnersConfig struct {
	path           string            `yaml:"-"`
	ConnectionURLS map[string]string `yaml:"conns"`
}

func (cfg *runnersConfig) parse() {
	if cfg.path == "" {
		slog.Error("Unable to get runners config path from environment")
		os.Exit(1)
	}

	confData, err := os.ReadFile(cfg.path)
	if err != nil {
		slog.Error("Unable to read runners config file", "error", err)
		os.Exit(1)
	}

	if err := yaml.Unmarshal(confData, cfg); err != nil {
		slog.Error("Unable to parse runners config file", "error", err)
		os.Exit(1)
	}
}

var RunnersConfig = runnersConfig{
	path: os.Getenv("RUNNERS_CONF_PATH"),
}
