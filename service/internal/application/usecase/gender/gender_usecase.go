package gender

import (
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/domain/entity"
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/domain/repository"
)

type GenderUsecase struct {
	genderRepo repository.GenderRepository
}

func NewGenderUsecase(genderRepo repository.GenderRepository) *GenderUsecase {
	return &GenderUsecase{genderRepo: genderRepo}
}

func (uc *GenderUsecase) FindByID(id int64) (*entity.Gender, error) {
	return uc.genderRepo.FindByID(id)
}
