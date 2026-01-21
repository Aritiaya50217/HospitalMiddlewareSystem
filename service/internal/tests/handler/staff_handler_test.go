package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/adapter/http/handler"
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/application/usecase/staff"
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/domain/entity"
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/tests/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateStaffHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

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
			return &entity.User{ID: id, RoleID: 2, HospitalID: 1}, nil
		},
		CreateFn: func(user *entity.User) error { return nil },
	}

	// Hospital repository mock
	hospitalRepo := &mocks.HospitalRepositoryMock{
		FindByNameFn: func(name string) (*entity.Hospital, error) {
			if name == "Bangkok Hospital" {
				return hospital, nil
			}
			return nil, errors.New("hospital not found")
		},
	}

	uc := staff.NewUsecaseStaff(userRepo, hospitalRepo)
	h := handler.NewUserHandler(uc)

	t.Run("Create Staff Handler", func(t *testing.T) {
		tests := []struct {
			name         string
			userID       any
			requestBody  any
			mockSetup    func()
			expectedCode int
			expectedMsg  string
		}{
			{
				name:   "success",
				userID: int64(1),
				requestBody: map[string]interface{}{
					"username": "staff1",
					"password": "staff1234",
					"hospital": "Bangkok Hospital",
				},
				mockSetup:    func() {},
				expectedCode: http.StatusCreated,
				expectedMsg:  "staff created successfully",
			},
			{
				name:   "forbidden non-admin",
				userID: int64(2),
				requestBody: map[string]interface{}{
					"username": "staff2",
					"password": "staff1234",
					"hospital": "Bangkok Hospital",
				},
				mockSetup:    func() {},
				expectedCode: http.StatusForbidden,
				expectedMsg:  "forbidden",
			},
			{
				name:   "hospital not found",
				userID: int64(1),
				requestBody: map[string]interface{}{
					"username": "staff3",
					"password": "staff1234",
					"hospital": "Unknown Hospital",
				},
				mockSetup: func() {
					hospitalRepo.FindByNameFn = func(name string) (*entity.Hospital, error) {
						return nil, errors.New("hospital not found")
					}
				},
				expectedCode: http.StatusInternalServerError,
				expectedMsg:  "hospital not found",
			},
			{
				name:   "create user error",
				userID: int64(1),
				requestBody: map[string]interface{}{
					"username": "staff4",
					"password": "staff1234",
					"hospital": "Bangkok Hospital",
				},
				mockSetup: func() {
					userRepo.CreateFn = func(user *entity.User) error { return errors.New("db error") }
				},
				expectedCode: http.StatusInternalServerError,
			},
			{
				name:         "missing user_id",
				userID:       nil,
				requestBody:  map[string]interface{}{"username": "staff5", "password": "123", "hospital": "Bangkok Hospital"},
				mockSetup:    func() {},
				expectedCode: http.StatusUnauthorized,
				expectedMsg:  "unauthorized: missing user_id",
			},
			{
				name:         "invalid JSON",
				userID:       int64(1),
				requestBody:  "invalid-json",
				mockSetup:    func() {},
				expectedCode: http.StatusBadRequest,
				expectedMsg:  "invalid request:",
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if tt.mockSetup != nil {
					tt.mockSetup()
				}

				var bodyReader *bytes.Reader
				switch v := tt.requestBody.(type) {
				case string:
					bodyReader = bytes.NewReader([]byte(v))
				default:
					b, _ := json.Marshal(v)
					bodyReader = bytes.NewReader(b)
				}

				req, _ := http.NewRequest("POST", "/staff/create", bodyReader)
				req.Header.Set("Content-Type", "application/json")

				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request = req
				if tt.userID != nil {
					c.Set("user_id", tt.userID)
				}

				h.CreateStaff(c)

				assert.Equal(t, tt.expectedCode, w.Code)
				var resp map[string]interface{}
				json.Unmarshal(w.Body.Bytes(), &resp)
				if tt.expectedCode == http.StatusCreated {
					assert.Equal(t, tt.expectedMsg, resp["message"])
				} else {
					assert.Contains(t, resp["error"], tt.expectedMsg)
				}
			})
		}
	})
}
