package create_staff

import (
	"errors"

	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/domain/entity"
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrForbidden = errors.New("forbidden")
)

type UsecaseCreate struct {
	userRepo     repository.UserRepository
	hospitalRepo repository.HospitalRepository
}

func NewUsecaseCreate(userRepo repository.UserRepository, hospitalRepo repository.HospitalRepository) *UsecaseCreate {
	return &UsecaseCreate{userRepo: userRepo, hospitalRepo: hospitalRepo}
}

func (uc *UsecaseCreate) Excute(id int64, req *CreateStaffRequest) error {
	// check role
	adminUser, err := uc.userRepo.FindByID(id)
	if err != nil {
		return errors.New("forbidden")
	}

	if adminUser.RoleID != 1 {
		return errors.New("forbidden")
	}

	// check hospital
	hospital, err := uc.hospitalRepo.FindByName(req.Hospital)
	if err != nil {
		return errors.New("hospital not found")
	}

	// hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// new member
	staff := &entity.User{
		Username:   req.Username,
		Password:   string(hashed),
		HospitalID: hospital.ID,
		RoleID:     2, // staff only
	}

	return uc.userRepo.Create(staff)
}
