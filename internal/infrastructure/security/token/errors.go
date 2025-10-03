package token

import (
	"errors"
	"fmt"
)

type ErrTokenNoRequiredData struct {
	missingField string
}

func (e ErrTokenNoRequiredData) Error() string {
	return fmt.Sprintf("token payload does not contain field %s", e.missingField)
}

var (
	ErrTokenExpired    = errors.New("token expired")
	ErrInvalidToken    = errors.New("invalid token")
	ErrTokenDataFormat = errors.New("incorrect token payload data format")
	ErrSigningToken    = errors.New("error signing token")
	ErrValidateToken   = errors.New("error validating token")
)
