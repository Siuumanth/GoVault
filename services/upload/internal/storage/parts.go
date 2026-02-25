package storage

import (
	"context"
	"log"
	"os"
	"strings"
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

	req, err := s.PresignClient.PresignUploadPart(
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
	// in dev, replace internal docker host with localhost for frontend use
	publicEndpoint := os.Getenv("AWS_PUBLIC_ENDPOINT")
	if publicEndpoint != "" {
		internalEndpoint := os.Getenv("AWS_ENDPOINT")
		url := strings.Replace(req.URL, internalEndpoint, publicEndpoint, 1)
		log.Printf("replacing %s with %s to get %s", internalEndpoint, publicEndpoint, url)
		return url, nil
	}

	log.Printf("presign upload part url is %s", req.URL)

	return req.URL, nil
}
