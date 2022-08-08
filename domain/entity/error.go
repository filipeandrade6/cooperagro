package entity

import "errors"

var (
	ErrNotFound      = errors.New("not found")
	ErrInvalidEntity = errors.New("invalid entity")
)
