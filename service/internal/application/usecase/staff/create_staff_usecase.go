package staff

import (
	"errors"

	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/domain/entity"
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrForbidden = errors.New("forbidden")
)

type UsecaseStaff struct {
	userRepo     repository.UserRepository
	hospitalRepo repository.HospitalRepository
}

func NewUsecaseStaff(userRepo repository.UserRepository, hospitalRepo repository.HospitalRepository) *UsecaseStaff {
	return &UsecaseStaff{userRepo: userRepo, hospitalRepo: hospitalRepo}
}

func (uc *UsecaseStaff) Excute(id int64, req *CreateStaffRequest) error {
	// check role
	adminUser, err := uc.userRepo.FindByID(id)
	if err != nil {
		return err
	}

	if adminUser.RoleID != 1 {
		return ErrForbidden
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

func (uc *UsecaseStaff) DeleteStaffByID(adminID, staffID int64) error {
	// check role
	adminUser, err := uc.userRepo.FindByID(adminID)
	if err != nil {
		return err
	}

	if adminUser.RoleID != 1 {
		return ErrForbidden
	}

	staff, err := uc.FindByID(staffID)
	if err != nil {
		return err
	}

	// check hospital
	hospital, err := uc.hospitalRepo.FindByID(staff.HospitalID)
	if err != nil {
		return err
	}

	if hospital.ID != adminUser.HospitalID {
		return errors.New("invalid hospital")
	}

	return uc.userRepo.Delete(staffID)
}

func (uc *UsecaseStaff) FindByID(id int64) (*entity.User, error) {
	user, err := uc.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
