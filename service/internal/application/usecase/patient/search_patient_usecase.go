package patient

import (
	"errors"

	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/domain/entity"
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/domain/repository"
)

var (
	ErrPatientNotFound     = errors.New("patient not found")
	ErrEmptySearchCriteria = errors.New("at least one search criteria is required")
	ErrInvalidDateOfBirth  = errors.New("invalid date_of_birth format (YYYY-MM-DD)")
)

type PatientUsecase struct {
	patientRepo repository.PatientRepository
}

func NewPatientUsecase(patientRepo repository.PatientRepository) *PatientUsecase {
	return &PatientUsecase{patientRepo: patientRepo}
}

func (uc *PatientUsecase) Search(hospitalID int64, nationalID, passportID, firstname, middlename, lastname, dateOfBirth, phoneNumber, email string) ([]entity.Patient, error) {

	if nationalID == "" && passportID == "" && firstname == "" &&
		middlename == "" && lastname == "" && dateOfBirth == "" &&
		phoneNumber == "" && email == "" {
		return nil, ErrEmptySearchCriteria
	}

	return uc.patientRepo.Search(hospitalID, nationalID, passportID, firstname, middlename, lastname, dateOfBirth, phoneNumber, email)
}

func (uc *PatientUsecase) SearchByID(id string, hospitalID int64) (*entity.Patient, error) {

	patient, err := uc.patientRepo.SearchByID(id, hospitalID)
	if err != nil {
		return nil, err
	}

	if patient == nil {
		return nil, ErrPatientNotFound
	}

	return patient, nil
}
