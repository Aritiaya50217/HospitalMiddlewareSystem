package usecase

import (
	"errors"
	"testing"

	create_staff "github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/application/usecase/staff"
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/domain/entity"
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCreateStaffUsecase(t *testing.T) {
	admin := &entity.User{ID: 1, RoleID: 1, HospitalID: 1}

	hospital := &entity.Hospital{ID: 1, Name: "Bangkok Hospital"}

	userRepo := &mocks.UserRepositoryMock{
		FindByIDFn: func(id int64) (*entity.User, error) {
			if id == 1 {
				return admin, nil
			}
			return nil, errors.New("not found")
		},
		CreateFn: func(user *entity.User) error {
			return nil
		},
	}

	hospitalRepo := &mocks.HospitalRepositoryMock{
		FindByNameFn: func(name string) (*entity.Hospital, error) {
			if name == "Bangkok Hospital" {
				return hospital, nil
			}
			return nil, errors.New("not found")
		},
	}

	uc := create_staff.NewUsecaseCreate(userRepo, hospitalRepo)

	t.Run("success", func(t *testing.T) {
		req := &create_staff.CreateStaffRequest{
			Username: "staff1",
			Password: "staff1234",
			Hospital: "Bangkok Hospital",
		}
		err := uc.Excute(admin.ID, req)
		assert.NoError(t, err)
	})

	t.Run("forbidden non-admin", func(t *testing.T) {
		req := &create_staff.CreateStaffRequest{
			Username: "staff2",
			Password: "staff1234",
			Hospital: "Bangkok Hospital",
		}

		nonAdmin := &entity.User{ID: 2, RoleID: 2}

		userRepo.FindByIDFn = func(id int64) (*entity.User, error) {
			return nonAdmin, nil
		}
		err := uc.Excute(nonAdmin.ID, req)
		assert.Equal(t, create_staff.ErrForbidden, err)
	})

	t.Run("hospital not found", func(t *testing.T) {
		req := &create_staff.CreateStaffRequest{
			Username: "staff3",
			Password: "staff1234",
			Hospital: "Unknown Hospital",
		}
		err := uc.Excute(admin.ID, req)
		assert.Error(t, err)
	})
}
