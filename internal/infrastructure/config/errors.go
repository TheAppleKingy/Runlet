package config

import (
	"errors"
)

var (
	ErrNoEnvUsername   = errors.New("POSTGRES_USER was not set at env")
	ErrNoEnvDbName     = errors.New("POSTGRES_DB was not set at env")
	ErrNoEnvDbPassword = errors.New("POSTGRES_PASSWORD was not set at env")
)

var (
	ErrNoRunnersConfig    = errors.New("RUNNERS_CONF_PATH was not set at env")
	ErrReadRunnersConfig  = errors.New("unable to read runners config file")
	ErrParseRunnersConfig = errors.New("unable to parse runners config file")
)

var (
	ErrNoEnvTokenExpTime = errors.New("JWT_TOKEN_EXPIRE_TIME was not set at env")
	ErrInvalidExpFormat  = errors.New("cannot parse int value from JWT_TOKEN_EXPIRE_TIME")
	ErrNoEnvSecret       = errors.New("SECRET_KEY was not set at env")
)

var (
	ErrNoEnvDebug         = errors.New("DEBUG was not set at env")
	ErrInvalidDebugFormat = errors.New("cannot parse bool from DEBUG")
)
