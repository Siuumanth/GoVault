package shared

import "errors"

var ErrChunkAlreadyExists error = errors.New("chunk already exists")
var ErrAcceptedAsync = errors.New("accepted for async processing")
var ErrPartAlreadyExists error = errors.New("part already exists")
