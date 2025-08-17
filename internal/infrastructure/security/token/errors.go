package token

import (
	"errors"
	"fmt"
)

type ErrValidateToken struct {
	missingField string
}

func (e ErrValidateToken) Error() string {
	return fmt.Sprintf("token payload does not contain field %s", e.missingField)
}

var (
	ErrTokenExpired    = errors.New("token expired")
	ErrInvalidToken    = errors.New("invalid token")
	ErrTokenDataFormat = errors.New("incorrect token payload data format")
)
