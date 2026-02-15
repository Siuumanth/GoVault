package storage

import (
	"context"
	"time"
)

type FileStorage interface {
	Copy(ctx context.Context, srcKey, dstKey string) error
	Delete(ctx context.Context, key string) error
	GenerateDownloadURL(ctx context.Context, key string, expiry time.Duration, fileName string, mimeType string) (string, error)
}
