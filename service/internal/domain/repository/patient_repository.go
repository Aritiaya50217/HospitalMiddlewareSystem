package repository

import "github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/domain/entity"

type PatientRepository interface {
	Search(id string, hospitalID int64, offset, limit int) ([]*entity.Patient, error)
}
