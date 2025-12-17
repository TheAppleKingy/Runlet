package config

import (
	"os"
	"strconv"
)

type appConfig struct {
	IsDebug bool
}

func (conf *appConfig) parse() error {
	dbg := os.Getenv("DEBUG")
	if dbg == "" {
		return ErrNoEnvDebug
	}
	isDebug, err := strconv.ParseBool(dbg)
	if err != nil {
		return ErrInvalidDebugFormat
	}
	conf.IsDebug = isDebug
	return nil
}
