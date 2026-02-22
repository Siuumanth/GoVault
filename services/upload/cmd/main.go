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

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

// Done: Make it raw bytes uplaod
// TODO: add sending recieved chunks info with every request
// Done: Delete session if fail or after upload
func main() {
	godotenv.Load()
	// ---------- DB ----------
	dbURL := os.Getenv("UPLOAD_POSTGRES_URL_DEV")
	db, err := database.Connect(dbURL)
	if err != nil {
		log.Fatal(err)
	}

	repos := repository.NewRegistryFromDB(db)
	// ---------- AWS / S3 ----------
	s3Client := getS3Client()
	bucket := os.Getenv("BUCKET_NAME")
	s3Storage := storage.NewS3Storage(s3Client, bucket)

	// ---------- Service & Client ----------
	fsURL := os.Getenv("GOVAULT_FILES_SERVICE_URL")
	if dbURL == "" || fsURL == "" || bucket == "" {
		log.Fatal("missing required env vars")
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

	log.Println("Upload service running on :9002")
	log.Fatal(server.ListenAndServe())
}

func getS3Client() *s3.Client {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return s3.NewFromConfig(cfg)
}
