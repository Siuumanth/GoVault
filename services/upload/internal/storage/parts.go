package storage

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (s *S3Storage) PresignUploadPart(
	ctx context.Context,
	key string,
	uploadID string,
	partNumber int32,
	expiry time.Duration,
) (string, error) {

	presigner := s3.NewPresignClient(s.Client)

	req, err := presigner.PresignUploadPart(
		ctx,
		&s3.UploadPartInput{
			Bucket:     aws.String(s.Bucket),
			Key:        aws.String(key),
			UploadId:   aws.String(uploadID),
			PartNumber: &partNumber,
		},
		s3.WithPresignExpires(expiry),
	)
	if err != nil {
		return "", err
	}

	return req.URL, nil
}
