package hospital

import (
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/domain/entity"
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/domain/repository"
)

type HospitalUsecase struct {
	repo repository.HospitalRepository
}

func NewHospitalUsecase(repo repository.HospitalRepository) *HospitalUsecase {
	return &HospitalUsecase{repo: repo}
}

func (uc *HospitalUsecase) FindByID(id int64) (*entity.Hospital, error) {

	hospital, err := uc.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return hospital, nil
}
