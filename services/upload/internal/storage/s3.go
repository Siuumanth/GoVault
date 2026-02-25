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
	Client        *s3.Client
	PresignClient *s3.PresignClient
	Bucket        string
}

func NewS3Storage(internalclient *s3.Client, presignClient *s3.Client, bucket string) *S3Storage {
	return &S3Storage{
		Client:        internalclient,
		PresignClient: s3.NewPresignClient(presignClient),
		Bucket:        bucket,
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
		Bucket:          aws.String(s.Bucket),
		Key:             aws.String(key),
		UploadId:        aws.String(uploadID),
		MultipartUpload: &types.CompletedMultipartUpload{Parts: parts},
	}

	result, err := s.Client.CompleteMultipartUpload(ctx, input)
	if err != nil {
		return "", fmt.Errorf("failed to complete multipart: %w", err)
	}

	if result.Location != nil {
		return *result.Location, nil
	}
	return key, nil // Fallback to key if location is nil
}

func (s *S3Storage) AbortMultipart(ctx context.Context, key string, uploadID string) error {
	_, err := s.Client.AbortMultipartUpload(ctx, &s3.AbortMultipartUploadInput{
		Bucket:   aws.String(s.Bucket),
		Key:      aws.String(key),
		UploadId: aws.String(uploadID),
	})
	return err
}
