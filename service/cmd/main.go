package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/adapter/http"
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/adapter/http/handler"
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/adapter/persistence/postgres"
	create_staff_usecase "github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/application/usecase/create_staff"
	login_usecase "github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/application/usecase/login"
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/infrastructure/database"
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/infrastructure/security"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	hash, _ := bcrypt.GenerateFromPassword([]byte("admin1234"), bcrypt.DefaultCost)
	fmt.Println("hash : ", string(hash))

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
	hospitalRepo := postgres.NewHospitalRepository(db)

	jwtService := security.NewJWTService(jwtSecret)
	loginUsecase := login_usecase.NewLoginUsecase(userReppo, authRepo, jwtService)
	userUsecase := create_staff_usecase.NewUsecaseCreate(userReppo, hospitalRepo)

	authHandler := handler.NewAuthHandler(loginUsecase)
	userHandler := handler.NewUserHandler(userUsecase)

	r := http.NewRouter(http.Router{
		AuthHandler: authHandler,
		UserHandler: userHandler,
		JWTService:  jwtService,
	})
	r.Run(":" + port)

}
