package http

import (
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/adapter/http/handler"
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/infrastructure/security"
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/middleware"
	"github.com/gin-gonic/gin"
)

type Router struct {
	AuthHandler *handler.AuthHandler
	UserHandler *handler.UserHandler
	JWTService  *security.JWTService
}

func NewRouter(router Router) *gin.Engine {
	r := gin.Default()

	r.POST("/staff/login", router.AuthHandler.Login)

	staff := r.Group("/staff")
	staff.Use(middleware.AuthMiddleware(router.JWTService))
	{
		staff.POST("/create", router.UserHandler.CreateStaff)
	}

	return r
}
