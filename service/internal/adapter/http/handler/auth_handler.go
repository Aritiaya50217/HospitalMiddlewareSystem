package handler

import (
	"net/http"

	usecase "github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/application/usecase/login"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	loginUsecase *usecase.LoginUsecase
}

func NewAuthHandler(loginUsecase *usecase.LoginUsecase) *AuthHandler {
	return &AuthHandler{loginUsecase: loginUsecase}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req usecase.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	res, err := h.loginUsecase.Login(&req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": res.AccessToken})
}
