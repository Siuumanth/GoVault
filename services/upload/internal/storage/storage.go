package storage

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type FileStorage interface {
	UploadFile(ctx context.Context, key string, localPath string) error
	// New Methods for Multipart
	InitiateMultipart(ctx context.Context, key string) (string, error)
	CompleteMultipart(ctx context.Context, key string, uploadID string, parts []types.CompletedPart) (string, error)
}
