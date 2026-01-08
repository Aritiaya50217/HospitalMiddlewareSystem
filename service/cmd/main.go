package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/adapter/http"
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/adapter/http/handler"
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/adapter/persistence/postgres"
	usecase "github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/application/usecase/login"
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/infrastructure/database"
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/infrastructure/security"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	hash, _ := bcrypt.GenerateFromPassword([]byte("admin1234"), bcrypt.DefaultCost)
	fmt.Println("hash : ",string(hash))

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET environment variable is not set")
	}

	// connect to database
	db := database.NewPostgresDB()

	userReppo := postgres.NewUserRepository(db)
	authRepo := postgres.NewAuthRepository(db)

	jwtService := security.NewJWTService(jwtSecret)
	loginUsecase := usecase.NewLoginUsecase(userReppo, authRepo, jwtService)

	authHandler := handler.NewAuthHandler(loginUsecase)

	r := http.NewRouter(authHandler)
	r.Run(":" + port)

}
