package repository

import "github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/domain/entity"

type GenderRepository interface {
	FindByID(id int64) (*entity.Gender, error)
}
