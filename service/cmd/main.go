package main

import (
	"log"
	"os"

	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/adapter/http"
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/adapter/http/handler"
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/adapter/persistence/postgres"
	auth "github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/application/usecase/auth"
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/application/usecase/gender"
	patient "github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/application/usecase/patient"
	staff "github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/application/usecase/staff"
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/infrastructure/database"
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/infrastructure/security"
)

func main() {
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
	patientRepo := postgres.NewPatientRepository(db)
	genderRepo := postgres.NewgenderRepository(db)

	jwtService := security.NewJWTService(jwtSecret)
	authUsecase := auth.NewLoginUsecase(userReppo, authRepo, jwtService)
	userUsecase := staff.NewUsecaseStaff(userReppo, hospitalRepo)
	patientUsecase := patient.NewPatientUsecase(patientRepo)
	genderUsecase := gender.NewGenderUsecase(genderRepo)

	authHandler := handler.NewAuthHandler(authUsecase)
	userHandler := handler.NewUserHandler(userUsecase)
	patientHandler := handler.NewPatientHandler(patientUsecase, genderUsecase)

	r := http.NewRouter(http.Router{
		AuthHandler:    authHandler,
		UserHandler:    userHandler,
		JWTService:     jwtService,
		PatientHandler: patientHandler,
	})
	r.Run(":" + port)

}
