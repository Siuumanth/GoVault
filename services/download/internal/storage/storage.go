package storage

import "context"

type FileStorage interface {
	GenerateDownloadURL(
		ctx context.Context,
		key string,
		expirySeconds int64,
	) (string, error)
}
