package storage

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type S3Storage struct {
	Client *s3.Client
	Bucket string
}

func NewS3Storage(client *s3.Client, bucket string) *S3Storage {
	return &S3Storage{
		Client: client,
		Bucket: bucket,
	}
}

// UploadFile reads from a file and puts the data into an object in a bucket.
func (s *S3Storage) UploadFile(ctx context.Context, objectKey string, localPath string) error {
	file, err := os.Open(localPath)
	if err != nil {
		return fmt.Errorf("open file %s: %w", localPath, err)

	}
	defer file.Close()

	_, err = s.Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(objectKey),
		Body:   file,
	})

	return err
}

func (s *S3Storage) InitiateMultipart(ctx context.Context, key string) (string, error) {
	output, err := s.Client.CreateMultipartUpload(ctx, &s3.CreateMultipartUploadInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return "", err
	}
	return *output.UploadId, nil
}

// Upload Part

func (s *S3Storage) CompleteMultipart(ctx context.Context, key string, uploadID string, parts []types.CompletedPart) (string, error) {
	input := &s3.CompleteMultipartUploadInput{
		Bucket:   aws.String(s.Bucket),
		Key:      aws.String(key),
		UploadId: aws.String(uploadID),
		MultipartUpload: &types.CompletedMultipartUpload{
			Parts: parts,
		},
	}

	result, err := s.Client.CompleteMultipartUpload(ctx, input)
	if err != nil {
		return "", fmt.Errorf("failed to complete multipart: %w", err)
	}

	// return the final location of the file
	return *result.Location, nil
}
