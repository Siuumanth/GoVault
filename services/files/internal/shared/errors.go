package shared

import "errors"

var (
	ErrRowExists    = errors.New("duplicate rows")
	ErrRowNotFound  = errors.New("Record not found")
	ErrUnauthorized = errors.New("unauthorized")
)
