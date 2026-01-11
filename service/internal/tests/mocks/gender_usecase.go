package mocks

import "github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/domain/entity"

type GenderUsecaseMock struct {
	FindByIDFn func(id int64) (*entity.Gender, error)
}

func (m *GenderUsecaseMock) FindByID(id int64) (*entity.Gender, error) {
	return m.FindByIDFn(id)
}
