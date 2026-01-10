package mocks

import (
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/domain/entity"
)

type PatientRepositoryMock struct {
	SearchPatientsFn func(id string, hospitalID int64, offset, limit int) ([]*entity.Patient, error)
	SearchByIDFn     func(id string, hospitalID int64) (*entity.Patient, error)
}

func (m *PatientRepositoryMock) SearchPatients(id string, hospitalID int64, offset, limit int) ([]*entity.Patient, error) {
	return m.SearchPatientsFn(id, hospitalID, offset, limit)
}

func (m *PatientRepositoryMock) SearchByID(id string, hospitalID int64) (*entity.Patient, error) {
	return m.SearchByIDFn(id, hospitalID)
}
