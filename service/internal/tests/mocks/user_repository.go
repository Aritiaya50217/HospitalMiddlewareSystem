package mocks

import "github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/domain/entity"

type UserRepositoryMock struct {
	FindByIDFn                  func(id int64) (*entity.User, error)
	FindByUserNameAndHospitalFn func(username, hospital string) (*entity.User, error)
	CreateFn                    func(user *entity.User) error
	DeleteFn                    func(id int64) error
}

func (m *UserRepositoryMock) FindByID(id int64) (*entity.User, error) {
	if m.FindByIDFn != nil {
		return m.FindByIDFn(id)
	}
	return nil, nil
}

func (m *UserRepositoryMock) FindByUserNameAndHospital(username, hospital string) (*entity.User, error) {
	return m.FindByUserNameAndHospitalFn(username, hospital)
}

func (m *UserRepositoryMock) Create(user *entity.User) error {
	if m.CreateFn != nil {
		return m.CreateFn(user)
	}
	return nil
}

func (m *UserRepositoryMock) Delete(id int64) error {
	return m.DeleteFn(id)
}
