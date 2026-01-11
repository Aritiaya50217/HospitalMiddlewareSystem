package auth

import (
	"errors"

	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/domain/entity"
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type LoginUsecase struct {
	UserRepo repository.UserRepository
	AuthRepo repository.AuthRepository
	JWT      JWTService
}

func NewLoginUsecase(userRepo repository.UserRepository, authRepo repository.AuthRepository, jwt JWTService) *LoginUsecase {
	return &LoginUsecase{
		UserRepo: userRepo,
		AuthRepo: authRepo,
		JWT:      jwt,
	}
}

func (u *LoginUsecase) Login(req *LoginRequest) (*LoginResponse, error) {
	user, err := u.UserRepo.FindByUserNameAndHospital(req.Username, req.Hospital)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid password")
	}

	token, expires, err := u.JWT.Generate(user)
	if err != nil {
		return nil, err
	}

	auth := &entity.Auth{
		UserID:    user.ID,
		Token:     token,
		ExpiredAt: expires,
	}

	if err := u.AuthRepo.Create(auth); err != nil {
		return nil, err
	}

	return &LoginResponse{AccessToken: token}, nil
}
