package usecase

import (
	"errors"
	"testing"
	"time"

	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/application/usecase/auth"
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/domain/entity"
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/tests/mocks"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestLoginUsecase(t *testing.T) {
	hashed, _ := bcrypt.GenerateFromPassword([]byte("staff1234"), bcrypt.DefaultCost)

	// mock
	user := &entity.User{
		ID:         2,
		Username:   "staff1",
		Password:   string(hashed),
		HospitalID: 1,
		RoleID:     2,
	}

	userRepo := &mocks.UserRepositoryMock{
		FindByUserNameAndHospitalFn: func(username, hospital string) (*entity.User, error) {
			if username == "staff1" && hospital == "Bangkok Hospital" {
				return user, nil
			}
			return nil, errors.New("user not found")
		},
		CreateFn: func(user *entity.User) error {
			return nil
		},

		FindByIDFn: func(id int64) (*entity.User, error) {
			return nil, nil
		},
	}

	authRepo := &mocks.AuthRepositoryMock{
		CreateFn: func(auth *entity.Auth) error {
			return nil
		},
		FindByTokenFn: func(token string) (*entity.Auth, error) {
			return nil, nil
		},
	}

	jwtSvc := &mocks.JWTServiceMock{
		GenerateFn: func(user *entity.User) (string, time.Time, error) {
			return "mock.jwt.token", time.Now().Add(24 * time.Hour), nil
		},
		ValidateFn: func(token string) (*entity.User, error) {
			return user, nil
		},
	}

	uc := login.NewLoginUsecase(userRepo, authRepo, jwtSvc)

	t.Run("success", func(t *testing.T) {
		resp, err := uc.Login(&login.LoginRequest{
			Username: "staff1",
			Password: "staff1234",
			Hospital: "Bangkok Hospital",
		})
		assert.NoError(t, err)
		assert.Equal(t, "mock.jwt.token", resp.AccessToken)
	})

	t.Run("invalid password", func(t *testing.T) {
		_, err := uc.Login(&login.LoginRequest{
			Username: "staff1",
			Password: "wrongpass",
			Hospital: "Bangkok Hospital",
		})
		assert.Error(t, err)
		assert.Equal(t, "invalid password", err.Error())
	})

	t.Run("user not found", func(t *testing.T) {
		_, err := uc.Login(&login.LoginRequest{
			Username: "notexist",
			Password: "staff1234",
			Hospital: "Bangkok Hospital",
		})
		assert.Error(t, err)
		assert.Equal(t, "user not found", err.Error())
	})
}
