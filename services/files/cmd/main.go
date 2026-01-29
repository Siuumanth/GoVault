package main

import (
	"context"
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

func main() {
	godotenv.Load()
	// ---------- DB ----------
	dbURL := os.Getenv("GOVAULT_POSTGRES_URL_DEV")

	db, err := database.Connect(dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repoReg := repository.NewPostgresRegistry(db)
	// ---------- AWS / S3 ----------
	s3Client := getS3Client()
	bucket := os.Getenv("BUCKET_NAME")
	s3Storage := storage.NewS3Storage(s3Client, bucket)

	// ---------- SERvice ----------
	s := service.NewServiceRegistry(repoReg, s3Storage)

	// ---------- HTTP ----------
	h := handler.NewHandlerRegistry(s)
	r := router.NewConfiguredChiRouter(h.Files, h.Shares, h.Shortcuts, h.Health)
	server := &http.Server{
		Addr:         ":9003",
		Handler:      r,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Println("Upload service running on :9003")
	log.Fatal(server.ListenAndServe())

}

func getS3Client() *s3.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	return s3.NewFromConfig(cfg)
}
