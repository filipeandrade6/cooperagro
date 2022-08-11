package entity

import "errors"

var (
	// ErrNotFound not found
	ErrNotFound = errors.New("not found")
	// ErrInvalidEntity invalid entity
	ErrInvalidEntity = errors.New("invalid entity")
	// ErrEntityAlreadyExists entity already exists
	ErrEntityAlreadyExists = errors.New("entity already exists")
)
