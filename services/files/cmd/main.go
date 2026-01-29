package main

import (
	"files/internal/database"
	"files/internal/repository"
	"files/internal/service"
	"log"
	"os"

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
	st := storage.NewStorage(repoReg)

	// ---------- SERvice ----------
	s := service.NewServiceRegistry(repoReg)

}
