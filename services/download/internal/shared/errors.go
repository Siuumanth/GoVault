package shared

import (
	"errors"
	"fmt"
)

var (
	ErrRowExists     = errors.New("duplicate rows")
	ErrRowNotFound   = errors.New("Record not found")
	ErrUnauthorized  = errors.New("unauthorized")
	ErrTooManyShares = fmt.Errorf("Maximum File Share limit is %d", MAX_SHARES)
)
