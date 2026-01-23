package storage

import (
	"context"
)

type FileStorage interface {
	Copy(ctx context.Context, srcKey, dstKey string) error
	Delete(ctx context.Context, key string) error
}
