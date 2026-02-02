package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"download/internal/database"
	"download/internal/handler"
	"download/internal/repository"
	"download/internal/router"
	"download/internal/service"
	"download/internal/storage"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	dbURL := os.Getenv("GOVAULT_POSTGRES_URL_DEV")
	db, err := database.Connect(dbURL)
	if err != nil {
		log.Fatal(err)
	}

	repos := repository.NewRegistryFromDB(db)

	awsCfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	s3Client := s3.NewFromConfig(awsCfg)
	bucket := os.Getenv("BUCKET_NAME")
	s3Storage := storage.NewS3Storage(s3Client, bucket)

	downloadService := service.NewDownloadService(repos, s3Storage)
	downloadHandler := handler.NewDownloadHandler(downloadService)

	r := chi.NewRouter()
	router.RegisterDownloadRoutes(r, downloadHandler)

	server := &http.Server{
		Addr:         ":9003",
		Handler:      r,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 0, // streaming
	}

	log.Println("Download service running on :9003")
	log.Fatal(server.ListenAndServe())
}
