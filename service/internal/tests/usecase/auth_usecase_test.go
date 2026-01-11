package usecase

import (
	"errors"
	"testing"
	"time"

	auth "github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/application/usecase/auth"
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/domain/entity"
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/tests/mocks"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestLoginUsecase(t *testing.T) {
	hashed, _ := bcrypt.GenerateFromPassword([]byte("staff1234"), bcrypt.DefaultCost)

	user := &entity.User{
		ID:         2,
		Username:   "staff1",
		Password:   string(hashed),
		HospitalID: 1,
		RoleID:     2,
	}

	t.Run("success", func(t *testing.T) {
		userRepo := &mocks.UserRepositoryMock{
			FindByUserNameAndHospitalFn: func(username, hospital string) (*entity.User, error) {
				return user, nil
			},
		}
		authRepo := &mocks.AuthRepositoryMock{
			CreateFn: func(auth *entity.Auth) error { return nil },
		}
		jwtSvc := &mocks.JWTServiceMock{
			GenerateFn: func(u *entity.User) (string, time.Time, error) {
				return "mock.jwt.token", time.Now().Add(24 * time.Hour), nil
			},
		}
		uc := auth.NewLoginUsecase(userRepo, authRepo, jwtSvc)

		resp, err := uc.Login(&auth.LoginRequest{
			Username: "staff1",
			Password: "staff1234",
			Hospital: "Bangkok Hospital",
		})
		assert.NoError(t, err)
		assert.Equal(t, "mock.jwt.token", resp.AccessToken)
	})

	t.Run("invalid password", func(t *testing.T) {
		userRepo := &mocks.UserRepositoryMock{
			FindByUserNameAndHospitalFn: func(username, hospital string) (*entity.User, error) {
				return user, nil
			},
		}
		authRepo := &mocks.AuthRepositoryMock{
			CreateFn: func(auth *entity.Auth) error { return nil },
		}
		jwtSvc := &mocks.JWTServiceMock{
			GenerateFn: func(u *entity.User) (string, time.Time, error) {
				return "mock.jwt.token", time.Now().Add(24 * time.Hour), nil
			},
		}
		uc := auth.NewLoginUsecase(userRepo, authRepo, jwtSvc)

		_, err := uc.Login(&auth.LoginRequest{
			Username: "staff1",
			Password: "wrongpass",
			Hospital: "Bangkok Hospital",
		})
		assert.Error(t, err)
		assert.Equal(t, "invalid password", err.Error())
	})

	t.Run("user not found", func(t *testing.T) {
		userRepo := &mocks.UserRepositoryMock{
			FindByUserNameAndHospitalFn: func(username, hospital string) (*entity.User, error) {
				return nil, errors.New("not found")
			},
		}
		authRepo := &mocks.AuthRepositoryMock{}
		jwtSvc := &mocks.JWTServiceMock{}
		uc := auth.NewLoginUsecase(userRepo, authRepo, jwtSvc)

		_, err := uc.Login(&auth.LoginRequest{
			Username: "notexist",
			Password: "staff1234",
			Hospital: "Bangkok Hospital",
		})
		assert.Error(t, err)
		assert.Equal(t, "user not found", err.Error())
	})

	t.Run("JWT generate error", func(t *testing.T) {
		userRepo := &mocks.UserRepositoryMock{
			FindByUserNameAndHospitalFn: func(username, hospital string) (*entity.User, error) {
				return user, nil
			},
		}
		authRepo := &mocks.AuthRepositoryMock{
			CreateFn: func(auth *entity.Auth) error { return nil },
		}
		jwtSvc := &mocks.JWTServiceMock{
			GenerateFn: func(u *entity.User) (string, time.Time, error) {
				return "", time.Time{}, errors.New("jwt error")
			},
		}
		uc := auth.NewLoginUsecase(userRepo, authRepo, jwtSvc)

		_, err := uc.Login(&auth.LoginRequest{
			Username: "staff1",
			Password: "staff1234",
			Hospital: "Bangkok Hospital",
		})
		assert.Error(t, err)
		assert.Equal(t, "jwt error", err.Error())
	})

	t.Run("AuthRepo create error", func(t *testing.T) {
		userRepo := &mocks.UserRepositoryMock{
			FindByUserNameAndHospitalFn: func(username, hospital string) (*entity.User, error) {
				return user, nil
			},
		}
		authRepo := &mocks.AuthRepositoryMock{
			CreateFn: func(auth *entity.Auth) error { return errors.New("db error") },
		}
		jwtSvc := &mocks.JWTServiceMock{
			GenerateFn: func(u *entity.User) (string, time.Time, error) {
				return "mock.jwt.token", time.Now().Add(24 * time.Hour), nil
			},
		}
		uc := auth.NewLoginUsecase(userRepo, authRepo, jwtSvc)

		_, err := uc.Login(&auth.LoginRequest{
			Username: "staff1",
			Password: "staff1234",
			Hospital: "Bangkok Hospital",
		})
		assert.Error(t, err)
		assert.Equal(t, "db error", err.Error())
	})
}
