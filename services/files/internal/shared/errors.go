package shared

import "errors"

var (
	ErrRowExists = errors.New("duplicate rows")
)
