package postgres

import (
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/domain/entity"
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/domain/repository"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindByID(id int64) (*entity.User, error) {
	var userModel UserModel
	if err := r.db.Preload("Hospital").Preload("Role").First(&userModel, id).Error; err != nil {
		return nil, err
	}

	user := &entity.User{
		ID:         userModel.ID,
		Username:   userModel.Username,
		Password:   userModel.Password,
		HospitalID: userModel.HospitalID,
		RoleID:     userModel.RoleID,
	}
	return user, nil
}

func (r *userRepository) FindByUserNameAndHospital(username, hospitalName string) (*entity.User, error) {
	var userModel UserModel
	if err := r.db.Preload("Hospital").
		Preload("Role").
		Joins("JOIN hospitals ON hospitals.id = users.hospital_id").
		Where("users.username = ? AND hospitals.name = ?", username, hospitalName).
		Order("users.id").
		First(&userModel).Error; err != nil {
		return nil, err
	}

	user := &entity.User{
		ID:         userModel.ID,
		Username:   userModel.Username,
		Password:   userModel.Password,
		HospitalID: userModel.HospitalID,
		RoleID:     userModel.RoleID,
		CreatedAt: userModel.CreatedAt,
		UpdatedAt: userModel.UpdatedAt,
	}

	return user, nil
}

func (r *userRepository) Create(user *entity.User) error {
	model := &UserModel{
		Username:   user.Username,
		Password:   user.Password,
		HospitalID: user.HospitalID,
		RoleID:     user.RoleID,
	}

	return r.db.Create(model).Error
}
