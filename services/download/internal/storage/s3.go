package storage

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Storage struct {
	Client    *s3.Client
	Presigner *s3.PresignClient
	Bucket    string
}

func NewS3Storage(client *s3.Client, bucket string) *S3Storage {
	return &S3Storage{
		Client:    client,
		Presigner: s3.NewPresignClient(client),
		Bucket:    bucket,
	}
}

func (s *S3Storage) GenerateDownloadURL(
	ctx context.Context,
	key string,
	expirySeconds int64,
) (string, error) {

	out, err := s.Presigner.PresignGetObject(
		ctx,
		&s3.GetObjectInput{
			Bucket: aws.String(s.Bucket),
			Key:    aws.String(key),
		},
		s3.WithPresignExpires(time.Duration(expirySeconds)*time.Second),
	)
	if err != nil {
		return "", err
	}

	return out.URL, nil
}
