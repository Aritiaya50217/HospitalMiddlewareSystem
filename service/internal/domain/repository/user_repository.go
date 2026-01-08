package repository

import "github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/domain/entity"

type UserRepository interface {
	FindByUserNameAndHospital(username, hospital string) (*entity.User, error)
}
