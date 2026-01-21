package handler

import (
	"net/http"
	"strconv"

	hospital "github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/application/usecase/hospital"
	staff "github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/application/usecase/staff"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	usecaseStaff    *staff.UsecaseStaff
	usecaseHospital *hospital.HospitalUsecase
}

func NewUserHandler(usecaseStaff *staff.UsecaseStaff) *UserHandler {
	return &UserHandler{usecaseStaff: usecaseStaff}
}

func (h *UserHandler) CreateStaff(c *gin.Context) {
	var req staff.CreateStaffRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request: " + err.Error()})
		return
	}

	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized: missing user_id"})
		return
	}

	userID, ok := userIDVal.(int64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error: user_id type mismatch"})
		return
	}

	if err := h.usecaseStaff.Excute(userID, &req); err != nil {
		switch err {
		case staff.ErrForbidden:
			c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "staff created successfully"})
}

func (h *UserHandler) DeleteStaff(c *gin.Context) {
	staffID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized: missing user_id"})
		return
	}

	userID, ok := userIDVal.(int64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error: user_id type mismatch"})
		return
	}

	if err := h.usecaseStaff.DeleteStaffByID(userID, staffID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "delete successfully"})
}
