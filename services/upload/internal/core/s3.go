package core

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Storage interface {
	UploadFile(ctx context.Context, key string, localPath string) error
	GenerateDownloadURL(ctx context.Context, key string) (string, error)
}

type S3Storage struct {
	Client *s3.Client
	Bucket string
}

// UploadFile reads from a file and puts the data into an object in a bucket.
func (s *S3Storage) UploadFile(ctx context.Context, key, localPath string) error {
	file, err := os.Open(localPath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = s.Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(key),
		Body:   file,
	})

	return err
}
