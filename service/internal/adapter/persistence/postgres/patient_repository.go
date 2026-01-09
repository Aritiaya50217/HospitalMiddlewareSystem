package postgres

import (
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/domain/entity"
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/domain/repository"
	"gorm.io/gorm"
)

type patientRepository struct {
	db *gorm.DB
}

func NewPatientRepository(db *gorm.DB) repository.PatientRepository {
	return &patientRepository{db: db}
}

func (r *patientRepository) FindByIDAndHospital(patientID int64, hospitalID int64) (*entity.Patient, error) {
	
	return nil, nil
}

func (r *patientRepository) FindByNationalIDAndHospital(nationalID string, hospitalID int64) (*entity.Patient, error) {
	return nil, nil
}

func (r *patientRepository) FindByPassportIDAndHospital(passportID string, hospitalID int64) (*entity.Patient, error) {
	return nil, nil
}
