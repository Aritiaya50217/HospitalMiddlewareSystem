package repository

import "github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/domain/entity"

type PatientRepository interface {
	Search(hospitalID int64, nationalID, passportID, firstname, middlename, lastname, dateOfBirth, phoneNumber, email string) ([]entity.Patient, error)
	SearchByID(id string, hospitalID int64) (*entity.Patient, error)
}
