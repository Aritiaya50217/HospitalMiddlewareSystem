package mocks

import "github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/domain/entity"

type GenderRepositoryMock struct {
	FindByIDFn func(id int64) (*entity.Gender, error)
}

func (m *GenderRepositoryMock) FindByID(id int64) (*entity.Gender, error) {
	return m.FindByIDFn(id)
}
