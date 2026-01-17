package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

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

func main() {
	godotenv.Load()
	// ---------- DB ----------
	dbURL := os.Getenv("POSTGRES_URL")
	db, err := database.Connect(dbURL)
	if err != nil {
		log.Fatal(err)
	}

	repos := repository.NewRegistryFromDB(db)
	// ---------- AWS / S3 ----------
	awsCfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	s3Client := s3.NewFromConfig(awsCfg)
	bucket := os.Getenv("BUCKET_NAME")
	s3Storage := storage.NewS3Storage(s3Client, bucket)

	// ---------- Service ----------
	uploadService := service.NewUploadService(repos, s3Storage)

	// ---------- Handler ----------
	uploadHandler := handler.NewUploadHandler(uploadService)

	// ---------- Router ----------
	r := chi.NewRouter()
	router.RegisterUploadRoutes(r, uploadHandler)

	// ---------- Server ----------
	server := &http.Server{
		Addr:         ":9002",
		Handler:      r,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Println("Upload service running on :9002")
	log.Fatal(server.ListenAndServe())
}
