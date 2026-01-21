package mocks

import "github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/domain/entity"

type HospitalRepositoryMock struct {
	FindByNameFn func(name string) (*entity.Hospital, error)
	FindByIDFn   func(id int64) (*entity.Hospital, error)
}

func (m *HospitalRepositoryMock) FindByName(name string) (*entity.Hospital, error) {
	return m.FindByNameFn(name)
}

func (m *HospitalRepositoryMock) FindByID(id int64) (*entity.Hospital, error) {
	return m.FindByIDFn(id)
}
