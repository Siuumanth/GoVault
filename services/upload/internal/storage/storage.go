package storage

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type FileStorage interface {
	UploadFile(ctx context.Context, key string, localPath string) error
	// New Methods for Multipart
	InitiateMultipart(ctx context.Context, key string) (string, error)
	PresignUploadPart(
		ctx context.Context,
		key string,
		uploadID string,
		partNumber int32,
		expiry time.Duration,
	) (string, error)
	CompleteMultipart(ctx context.Context, key string, uploadID string, parts []types.CompletedPart) (string, error)
	AbortMultipart(ctx context.Context, key string, uploadID string) error
}
