package mocks

import (
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/domain/entity"
)

type PatientRepositoryMock struct {
	SearchByIDFn func(id string, hospitalID int64) (*entity.Patient, error)
}

func (m *PatientRepositoryMock) SearchByID(id string, hospitalID int64) (*entity.Patient, error) {
	return m.SearchByIDFn(id, hospitalID)
}
