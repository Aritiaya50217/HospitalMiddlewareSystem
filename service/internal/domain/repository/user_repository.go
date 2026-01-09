package repository

import "github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/domain/entity"

type UserRepository interface {
	FindByID(id int64) (*entity.User, error)
	FindByUserNameAndHospital(username, hospital string) (*entity.User, error)
	Create(user *entity.User) error
}
