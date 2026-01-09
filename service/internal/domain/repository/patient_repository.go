package repository

import "github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/domain/entity"

type PatientRepository interface {
	FindByIDAndHospital(patientID int64, hospitalID int64) (*entity.Patient, error)
	FindByNationalIDAndHospital(nationalID string, hospitalID int64) (*entity.Patient, error)
	FindByPassportIDAndHospital(passportID string, hospitalID int64) (*entity.Patient, error)
}
