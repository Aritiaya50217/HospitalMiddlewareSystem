package handler

import (
	"net/http"
	"strconv"

	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/application/usecase/gender"
	patient "github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/application/usecase/patient"
	"github.com/gin-gonic/gin"
)

type PatientHandler struct {
	patientUsecase *patient.PatientUsecase
	genderUsecase  *gender.GenderUsecase
}

func NewPatientHandler(patientUsecase *patient.PatientUsecase, genderUsecase *gender.GenderUsecase) *PatientHandler {
	return &PatientHandler{patientUsecase: patientUsecase, genderUsecase: genderUsecase}
}

func (h *PatientHandler) Search(c *gin.Context) {
	id := c.Param("id")

	hospitalID, exists := c.Get("hospital_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	// optional pagination
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	patients, err := h.patientUsecase.Search(id, hospitalID.(int64), offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	results := make([]map[string]interface{}, 0)
	for _, patient := range patients {
		gender, err := h.genderUsecase.FindByID(int64(patient.GenderID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		results = append(results, map[string]interface{}{
			"first_name_th":  patient.FirstNameTH,
			"middle_name_th": patient.MiddleNameTH,
			"last_name_th":   patient.LastNameTH,
			"first_name_en":  patient.FirstNameEN,
			"middle_name_en": patient.MiddleNameEN,
			"last_name_en":   patient.LastNameEN,
			"date_of_birth":  patient.DateOfBirth,
			"patient_hn":     patient.PatientHN,
			"passport_id":    patient.PassportID,
			"phone_number":   patient.PhoneNumber,
			"email":          patient.Email,
			"gender":         gender.Name,
		})
	}

	c.JSON(http.StatusOK, gin.H{"result": results})

}
