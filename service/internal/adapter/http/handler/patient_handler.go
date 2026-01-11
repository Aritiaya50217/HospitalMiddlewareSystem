package handler

import (
	"net/http"

	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/application/usecase/gender"
	patient "github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/application/usecase/patient"
	search_patient "github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/application/usecase/patient"
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
	var req search_patient.SearchPatientRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid query parameters"})
		return
	}

	hospitalID, ok := c.Get("hospital_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	patients, err := h.patientUsecase.Search(
		hospitalID.(int64),
		req.NationalID,
		req.PassportID,
		req.FirstName,
		req.MiddleName,
		req.LastName,
		req.DateOfBirth,
		req.PhoneNumber,
		req.Email,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var results []search_patient.PatientResponse
	for _, patient := range patients {
		gender, err := h.genderUsecase.FindByID(int64(patient.GenderID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		results = append(results, search_patient.PatientResponse{
			FirstNameTH:  patient.FirstNameTH,
			MiddleNameTH: patient.MiddleNameTH,
			LastNameTH:   patient.LastNameTH,
			FirstNameEN:  patient.FirstNameEN,
			MiddleNameEN: patient.MiddleNameEN,
			LastNameEN:   patient.LastNameEN,
			DateOfBirth:  patient.DateOfBirth,
			NationalID:   patient.NationalID,
			PatientHN:    patient.PatientHN,
			PassportID:   patient.PassportID,
			PhoneNumber:  patient.PhoneNumber,
			Email:        patient.Email,
			Gender:       gender.Abbreviation,
		})
	}

	c.JSON(http.StatusOK, gin.H{"patients": results})
}

func (h *PatientHandler) SearchByID(c *gin.Context) {
	id := c.Param("id")

	hospitalID, exists := c.Get("hospital_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	patient, err := h.patientUsecase.SearchByID(id, hospitalID.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	gender, err := h.genderUsecase.FindByID(int64(patient.GenderID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	results := search_patient.PatientResponse{
		FirstNameTH:  patient.FirstNameTH,
		MiddleNameTH: patient.MiddleNameTH,
		LastNameTH:   patient.LastNameTH,
		FirstNameEN:  patient.FirstNameEN,
		MiddleNameEN: patient.MiddleNameEN,
		LastNameEN:   patient.LastNameEN,
		DateOfBirth:  patient.DateOfBirth,
		NationalID:   patient.NationalID,
		PatientHN:    patient.PatientHN,
		PassportID:   patient.PassportID,
		PhoneNumber:  patient.PhoneNumber,
		Email:        patient.Email,
		Gender:       gender.Abbreviation,
	}

	c.JSON(http.StatusOK, gin.H{"patient": results})

}
