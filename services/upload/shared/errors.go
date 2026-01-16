package shared

import "errors"

var ErrChunkAlreadyExists error = errors.New("duplicate")
