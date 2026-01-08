package http

import (
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/adapter/http/handler"
	"github.com/gin-gonic/gin"
)

func NewRouter(authHandler *handler.AuthHandler) *gin.Engine {
	r := gin.Default()

	r.POST("/staff/login", authHandler.Login)

	return r
}
