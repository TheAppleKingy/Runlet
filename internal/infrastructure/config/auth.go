package config

import (
	"os"
	"strconv"
)

type authConfig struct {
	Secret          string
	TokenExpireTime int
}

func (conf *authConfig) parse() error {
	exp := os.Getenv("JWT_TOKEN_EXPIRE_TIME")
	if exp == "" {
		return ErrNoEnvTokenExpTime
	}
	s := os.Getenv("SECRET_KEY")
	if s == "" {
		return ErrNoEnvSecret
	}
	e, err := strconv.Atoi(exp)
	if err != nil {
		return ErrInvalidExpFormat
	}
	conf.TokenExpireTime = e
	conf.Secret = s
	return nil
}
