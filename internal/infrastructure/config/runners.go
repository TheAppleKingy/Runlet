package config

import (
	"Runlet/internal/pkg/errs"
	"os"

	"gopkg.in/yaml.v3"
)

type runnersConfig struct {
	ConnectionURLS map[string]string `yaml:"conns"`
}

func (cfg *runnersConfig) parse() error {
	path := os.Getenv("RUNNERS_CONF_PATH")
	if path == "" {
		return ErrNoRunnersConfig
	}
	confData, err := os.ReadFile(path)
	if err != nil {
		return errs.WrapErrors(ErrReadRunnersConfig, err)
	}

	if err := yaml.Unmarshal(confData, cfg); err != nil {
		return errs.WrapErrors(ErrParseRunnersConfig, err)
	}
	return nil
}

var RunnersConfig = runnersConfig{}
