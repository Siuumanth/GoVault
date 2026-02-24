package main

import (
	"context"
	"files/internal/clients"
	"files/internal/database"
	"files/internal/handler"
	"files/internal/repository"
	"files/internal/router"
	"files/internal/service"
	"files/internal/storage"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
)

// TODO: add graceful shutdown
func main() {
	godotenv.Load()
	// ---------- DB ----------
	dbURL := os.Getenv("FILES_POSTGRES_URL_DEV")
	if dbURL == "" {
		log.Fatal("Database URL is not set")
	}

	db, err := database.Connect(dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repoRegistry := repository.NewPostgresRegistry(db)
	// ---------- AWS / S3 ----------
	s3Client := getS3Client()
	bucket := os.Getenv("BUCKET_NAME")
	s3Storage := storage.NewS3Storage(s3Client, bucket)

	// ---------- SERvice ----------
	authServiceURL := os.Getenv("GOVAULT_AUTH_SERVICE_URL")
	if authServiceURL == "" {
		log.Fatal("Auth service URL is not set")
	}
	authClient := clients.NewAuthClient(authServiceURL)
	s := service.NewServiceRegistry(repoRegistry, s3Storage, authClient)

	// ---------- HTTP ----------
	h := handler.NewHandlerRegistry(s)
	filesRouter := router.NewConfiguredChiRouter(h.Files, h.Shares, h.Shortcuts, h.Health)
	server := &http.Server{
		Addr:         ":9003",
		Handler:      filesRouter,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Println("Files service running on :9003")
	log.Fatal(server.ListenAndServe())

}

func getS3Client() *s3.Client {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// if endpoint is set, use it (MinIO in dev), otherwise use real S3
	endpoint := os.Getenv("AWS_ENDPOINT")
	if endpoint != "" {
		return s3.NewFromConfig(cfg, func(o *s3.Options) {
			o.BaseEndpoint = &endpoint
			o.UsePathStyle = true // MinIO requires path style
		})
	}

	return s3.NewFromConfig(cfg)
}
