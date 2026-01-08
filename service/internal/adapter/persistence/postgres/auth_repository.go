package postgres

import (
	"errors"

	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/domain/entity"
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/domain/repository"
	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) repository.AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) Create(auth *entity.Auth) error {
	return r.db.Create(auth).Error
}

func (r *authRepository) FindByToken(token string) (*entity.Auth, error) {
	var auth entity.Auth
	err := r.db.Preload("User").Where("token = ?", token).First(&auth).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &auth, nil
}
