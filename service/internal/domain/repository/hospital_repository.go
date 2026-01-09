package repository

import "github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/domain/entity"

type HospitalRepository interface {
	FindByName(name string) (*entity.Hospital, error)
}
