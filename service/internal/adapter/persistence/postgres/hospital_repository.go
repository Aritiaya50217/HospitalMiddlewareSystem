package postgres

import (
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/domain/entity"
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/domain/repository"
	"gorm.io/gorm"
)

type hospitalRepository struct {
	db *gorm.DB
}

func NewHospitalRepository(db *gorm.DB) repository.HospitalRepository {
	return &hospitalRepository{db: db}
}

func (r *hospitalRepository) FindByName(name string) (*entity.Hospital, error) {
	var hospital HospitalModel
	if err := r.db.Where("name = ? ", name).First(&hospital).Error; err != nil {
		return nil, err
	}

	return &entity.Hospital{ID: hospital.ID, Name: hospital.Name}, nil
}
