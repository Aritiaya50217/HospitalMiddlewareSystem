package mocks

import (
	"time"

	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/domain/entity"
)

type JWTServiceMock struct {
	GenerateFn func(user *entity.User) (string, time.Time, error)
	ValidateFn func(token string) (*entity.User, error)
}

func (m *JWTServiceMock) Generate(user *entity.User) (string, time.Time, error) {
	return m.GenerateFn(user)
}

func (m *JWTServiceMock) Validate(token string) (*entity.User, error) {
	return m.ValidateFn(token)
}
