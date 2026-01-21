package usecase

import (
	"errors"
	"testing"

	staff "github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/application/usecase/staff"
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/domain/entity"
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCreateStaffUsecase(t *testing.T) {
	// Mock admin user
	admin := &entity.User{ID: 1, RoleID: 1, HospitalID: 1}

	// Mock hospital
	hospital := &entity.Hospital{ID: 1, Name: "Bangkok Hospital"}

	// User repository mock
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

	// Hospital repository mock
	hospitalRepo := &mocks.HospitalRepositoryMock{
		FindByNameFn: func(name string) (*entity.Hospital, error) {
			if name == "Bangkok Hospital" {
				return hospital, nil
			}
			return nil, errors.New("not found")
		},
	}

	uc := staff.NewUsecaseStaff(userRepo, hospitalRepo)

	t.Run("success", func(t *testing.T) {
		req := &staff.CreateStaffRequest{
			Username: "staff1",
			Password: "staff1234",
			Hospital: "Bangkok Hospital",
		}
		err := uc.Excute(admin.ID, req)
		assert.NoError(t, err)
	})

	t.Run("forbidden non-admin", func(t *testing.T) {
		nonAdmin := &entity.User{ID: 2, RoleID: 2, HospitalID: 1}
		userRepo.FindByIDFn = func(id int64) (*entity.User, error) {
			return nonAdmin, nil
		}

		req := &staff.CreateStaffRequest{
			Username: "staff2",
			Password: "staff1234",
			Hospital: "Bangkok Hospital",
		}
		err := uc.Excute(nonAdmin.ID, req)
		assert.Equal(t, staff.ErrForbidden, err)
	})

	t.Run("hospital not found", func(t *testing.T) {
		// restore admin
		userRepo.FindByIDFn = func(id int64) (*entity.User, error) { return admin, nil }

		req := &staff.CreateStaffRequest{
			Username: "staff3",
			Password: "staff1234",
			Hospital: "Unknown Hospital",
		}
		err := uc.Excute(admin.ID, req)
		assert.Error(t, err)
		assert.Equal(t, "hospital not found", err.Error())
	})

	t.Run("create user error", func(t *testing.T) {
		// restore admin & hospital
		userRepo.FindByIDFn = func(id int64) (*entity.User, error) { return admin, nil }
		hospitalRepo.FindByNameFn = func(name string) (*entity.Hospital, error) { return hospital, nil }

		// Force CreateFn to return error
		userRepo.CreateFn = func(user *entity.User) error { return errors.New("db error") }

		req := &staff.CreateStaffRequest{
			Username: "staff4",
			Password: "staff1234",
			Hospital: "Bangkok Hospital",
		}
		err := uc.Excute(admin.ID, req)
		assert.Error(t, err)
		assert.Equal(t, "db error", err.Error())
	})
}

// func TestDeleteStaffUsecase(t *testing.T) {
// 	// Mock admin user
// 	admin := &entity.User{ID: 1, RoleID: 1, HospitalID: 1}

// 	// Mock hospital
// 	hospital := &entity.Hospital{ID: 1, Name: "Bangkok Hospital"}

// 	// User repository mock
// 	userRepo := &mocks.UserRepositoryMock{
// 		FindByIDFn: func(id int64) (*entity.User, error) {
// 			if id == 1 {
// 				return admin, nil
// 			}
// 			return nil, errors.New("not found")
// 		},
// 		DeleteFn: func(id int64) error {
// 			return nil
// 		},
// 	}

// 	// Hospital repository mock
// 	hospitalRepo := &mocks.HospitalRepositoryMock{
// 		FindByNameFn: func(name string) (*entity.Hospital, error) {
// 			if name == "Bangkok Hospital" {
// 				return hospital, nil
// 			}
// 			return nil, errors.New("not found")
// 		},
// 	}

// 	uc := staff.NewUsecaseStaff(userRepo, hospitalRepo)

// 	t.Run("success", func(t *testing.T) {
// 		err := uc.DeleteStaffByID(1, 2)
// 		assert.NoError(t, err)
// 	})
// }
