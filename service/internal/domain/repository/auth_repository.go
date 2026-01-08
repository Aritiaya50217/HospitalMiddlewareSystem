package repository

import "github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/domain/entity"

type AuthRepository interface {
	Create(auth *entity.Auth) error
	FindByToken(token string) (*entity.Auth, error)
}
