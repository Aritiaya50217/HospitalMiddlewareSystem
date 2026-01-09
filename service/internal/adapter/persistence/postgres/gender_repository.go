package postgres

import (
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/domain/entity"
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/domain/repository"
	"gorm.io/gorm"
)

type genderRepository struct {
	db *gorm.DB
}

func NewgenderRepository(db *gorm.DB) repository.GenderRepository {
	return &genderRepository{db: db}
}

func (r *genderRepository) FindByID(id int64) (*entity.Gender, error) {
	var gender *entity.Gender

	if err := r.db.Where("id = ? ", id).First(&gender).Error; err != nil {
		return nil, err
	}
	return gender, nil
}
