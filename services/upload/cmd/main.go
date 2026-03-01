package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"upload/internal/clients"
	"upload/internal/database"
	"upload/internal/handler"
	"upload/internal/repository"
	"upload/internal/router"
	"upload/internal/service"
	"upload/internal/storage"
	zlog "upload/pkg/zlog"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	// development level logger
	zlog.Init()
	defer zlog.Sync()

	godotenv.Load()
	// ---------- DB ----------
	dbURL := os.Getenv("UPLOAD_POSTGRES_URL_DEV")
	db, err := database.Connect(dbURL)
	if err != nil {
		zlog.L.Error("Error connecting to DB ", zap.Error(err))
	}

	repos := repository.NewRegistryFromDB(db)
	// ---------- AWS / S3 ----------
	internalS3Client, presignBaseClient := getS3Clients()
	bucket := os.Getenv("BUCKET_NAME")
	s3Storage := storage.NewS3Storage(internalS3Client, presignBaseClient, bucket)

	// ---------- Service & Client ----------
	fsURL := os.Getenv("GOVAULT_FILES_SERVICE_URL")
	if dbURL == "" || fsURL == "" || bucket == "" {
		zlog.L.Error("missing required env vars")
	}

	fileClient := clients.NewFileClient(fsURL)
	sr := service.NewServiceRegistry(repos, s3Storage, fileClient)

	// ---------- Handler ----------
	uploadHandler := handler.NewUploadHandler(sr.Proxy, sr.Multipart)

	// ---------- Router ----------
	mainRouter := chi.NewRouter()
	router.RegisterUploadRoutes(mainRouter, uploadHandler)

	// ---------- Server ----------
	server := &http.Server{
		Addr:         ":9002",
		Handler:      mainRouter,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	zlog.L.Info("Upload service running on :9002")
	log.Fatal(server.ListenAndServe())
}

// below logic is to seamlessly switch from prod to de

func getS3Clients() (*s3.Client, *s3.Client) {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// internal client for PutObject, CreateMultipartUpload etc
	// below lines for MinIO: if AWS_ENDP isnt theres its s3, if ter its miniIO
	// if endpoint is set, use it (MinIO in dev), otherwise use real S3

	internalEndpoint := os.Getenv("AWS_ENDPOINT")
	internalClient := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = &internalEndpoint
		o.UsePathStyle = true
	})

	publicEndpoint := os.Getenv("AWS_PUBLIC_ENDPOINT")
	if publicEndpoint == "" {
		publicEndpoint = internalEndpoint
	}
	presignBaseClient := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = &publicEndpoint
		o.UsePathStyle = true
	})

	return internalClient, presignBaseClient
}
