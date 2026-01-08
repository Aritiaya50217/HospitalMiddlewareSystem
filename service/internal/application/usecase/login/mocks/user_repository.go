package mocks

import "github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/domain/entity"

type UserRepositoryMock struct {
	FindByUserNameAndHospitalFn func(username, hospital string) (*entity.User, error)
}

func (m *UserRepositoryMock) FindByUserNameAndHospital(username, hospital string) (*entity.User, error) {
	return m.FindByUserNameAndHospitalFn(username, hospital)
}
