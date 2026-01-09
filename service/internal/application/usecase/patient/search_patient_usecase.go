package search_patient

import (
	"errors"

	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/domain/entity"
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/domain/repository"
)

var ErrPatientNotFound = errors.New("patient not found")

type PatientUsecase struct {
	patientRepo repository.PatientRepository
}

func NewPatientUsecase(patientRepo repository.PatientRepository) *PatientUsecase {
	return &PatientUsecase{patientRepo: patientRepo}
}

func (uc *PatientUsecase) Search(id string, hospitalID int64, offset, limit int) ([]*entity.Patient, error) {
	if limit <= 0 || limit > 100 {
		limit = 20
	}
	return uc.patientRepo.Search(id, hospitalID, offset, limit)
}
