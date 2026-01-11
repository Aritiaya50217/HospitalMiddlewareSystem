package mocks

import (
	"errors"

	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/domain/entity"
)

type PatientUsecaseMock struct {
	SearchByIDFn func(id string, hospitalID int64) (*entity.Patient, error)
}

func (m *PatientUsecaseMock) SearchByID(id string, hospitalID int64) (*entity.Patient, error) {
	if m.SearchByIDFn != nil {
		return m.SearchByIDFn(id, hospitalID)
	}
	return nil, errors.New("not umplemented")
}
