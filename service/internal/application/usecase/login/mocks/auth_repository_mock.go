package mocks

import "github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/domain/entity"

type AuthRepositoryMock struct {
	CreateFn      func(auth *entity.Auth) error
	FindByTokenFn func(token string) (*entity.Auth, error)
	Err           error
}

func (m *AuthRepositoryMock) Create(auth *entity.Auth) error {
	return m.CreateFn(auth)
}

func (m *AuthRepositoryMock) FindByToken(token string) (*entity.Auth, error) {
	return m.FindByTokenFn(token)
}
