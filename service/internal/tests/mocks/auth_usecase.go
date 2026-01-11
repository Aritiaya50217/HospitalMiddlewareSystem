package mocks

import (
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/application/usecase/auth"
)

type LoginUsecaseMock struct {
	LoginFn func(req *auth.LoginRequest) (*auth.LoginResponse, error)
}

func (m *LoginUsecaseMock) Login(req *auth.LoginRequest) (*auth.LoginResponse, error) {
	return m.LoginFn(req)
}
