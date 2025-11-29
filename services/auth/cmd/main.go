package main

import (
	"auth/internal/dao"
	"auth/internal/database"
	"auth/internal/handler"
	"auth/internal/router"
	"auth/internal/service"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

/*
Here we have to define all db, dao, service handlers, etc and connect them all
*/
func main() {
	godotenv.Load() // loads .env from root
	dbURL := os.Getenv("POSTGRES_URL")
	db, err := database.Connect(dbURL)
	if err != nil {
		panic(err)
	}

	authDao := dao.NewPostgresUserDAO(db)
	authService := service.NewPGAuthService(authDao)

	authHandler := handler.NewAuthHandler(authService)

	userRouter := router.NewRouter(authHandler)

	http.ListenAndServe(":8080", userRouter)
}
