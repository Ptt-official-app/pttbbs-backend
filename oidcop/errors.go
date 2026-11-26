package oidcop

import "errors"

var (
	ErrInvalidClient = errors.New("invalid client")
	ErrInvalidToken  = errors.New("invalid token")

	ErrInvalidSigningKey = errors.New("invalid signing key")
)
