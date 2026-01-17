package storage

import (
	"context"
)

type FileStorage interface {
	UploadFile(ctx context.Context, key string, localPath string) error
	//	GenerateDownloadURL(ctx context.Context, key string) (string, error)
}
