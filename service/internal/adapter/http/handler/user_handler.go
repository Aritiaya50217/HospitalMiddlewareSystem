package handler

import (
	"net/http"

	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/application/usecase/create_staff"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	createStaff *create_staff.UsecaseCreate
}

func NewUserHandler(createStaff *create_staff.UsecaseCreate) *UserHandler {
	return &UserHandler{createStaff: createStaff}
}

func (h *UserHandler) CreateStaff(c *gin.Context) {
	var req create_staff.CreateStaffRequest

	// Bind JSON safely
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request: " + err.Error()})
		return
	}

	// Get user_id from middleware context
	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized: missing user_id"})
		return
	}

	// Type assertion
	userID, ok := userIDVal.(int64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error: user_id type mismatch"})
		return
	}

	// Execute usecase
	if err := h.createStaff.Excute(userID, &req); err != nil {
		switch err {
		case create_staff.ErrForbidden:
			c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "staff created successfully"})
}
