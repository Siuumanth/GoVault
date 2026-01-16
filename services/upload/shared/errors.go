package shared

import "errors"

var ErrChunkAlreadyExists error = errors.New("chunk already exists")
var ErrAcceptedAsync = errors.New("accepted for async processing")
