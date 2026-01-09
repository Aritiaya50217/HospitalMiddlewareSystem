package search_patient

import (
	"errors"

	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/domain/entity"
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/domain/repository"
)

var ErrPatientNotFound = errors.New("patient not found")

type UsecaseSearch struct {
	patientRepo repository.PatientRepository
}

func NewUsecaseSearch(patientRepo repository.PatientRepository) *UsecaseSearch {
	return &UsecaseSearch{patientRepo: patientRepo}
}

func (uc *UsecaseSearch) Execute(hospitalID int64, req SearchPatientRequest) (*entity.Patient, error) {
	switch {
	case req.PatientID != nil:
		return uc.patientRepo.FindByIDAndHospital(*req.PatientID, hospitalID)

	case req.NationalID != nil:
		return uc.patientRepo.FindByNationalIDAndHospital(*req.NationalID, hospitalID)

	case req.PassportID != nil:
		return uc.patientRepo.FindByPassportIDAndHospital(*req.PassportID, hospitalID)
	
	default:
		return nil, ErrPatientNotFound
	}
}
