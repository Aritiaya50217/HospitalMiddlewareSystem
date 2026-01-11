package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/adapter/http/handler"
	auth "github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/application/usecase/auth"
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/domain/entity"
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/tests/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestLoginHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	hashed, _ := bcrypt.GenerateFromPassword([]byte("staff1234"), bcrypt.DefaultCost)

	user := &entity.User{
		ID:         1,
		Username:   "staff1",
		Password:   string(hashed),
		HospitalID: 1,
		RoleID:     2,
	}

	// Mock repository และ JWT
	userRepo := &mocks.UserRepositoryMock{
		FindByUserNameAndHospitalFn: func(username, hospital string) (*entity.User, error) {
			if username == "staff1" && hospital == "Bangkok Hospital" {
				return user, nil
			}
			return nil, errors.New("user not found")
		},
	}

	authRepo := &mocks.AuthRepositoryMock{
		CreateFn: func(auth *entity.Auth) error { return nil },
	}

	jwtSvc := &mocks.JWTServiceMock{
		GenerateFn: func(u *entity.User) (string, time.Time, error) {
			if u.Username == "jwtfail" {
				return "", time.Time{}, errors.New("jwt error")
			}
			return "mock.jwt.token", time.Now().Add(24 * time.Hour), nil
		},
	}

	// สร้าง Usecase ตัวจริงด้วย mock repository
	loginUsecase := auth.NewLoginUsecase(userRepo, authRepo, jwtSvc)
	h := handler.NewAuthHandler(loginUsecase)

	tests := []struct {
		name         string
		body         string
		expectedCode int
		expectedKey  string
		expectedVal  string
	}{
		{
			name:         "success login",
			body:         `{"username":"staff1","password":"staff1234","hospital":"Bangkok Hospital"}`,
			expectedCode: http.StatusOK,
			expectedKey:  "token",
			expectedVal:  "mock.jwt.token",
		},
		{
			name:         "invalid password",
			body:         `{"username":"staff1","password":"wrongpass","hospital":"Bangkok Hospital"}`,
			expectedCode: http.StatusUnauthorized,
			expectedKey:  "error",
			expectedVal:  "invalid password",
		},
		{
			name:         "user not found",
			body:         `{"username":"notexist","password":"staff1234","hospital":"Bangkok Hospital"}`,
			expectedCode: http.StatusUnauthorized,
			expectedKey:  "error",
			expectedVal:  "user not found",
		},
		{
			name:         "invalid JSON",
			body:         `invalid-json`,
			expectedCode: http.StatusBadRequest,
			expectedKey:  "error",
			expectedVal:  "invalid request",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request, _ = http.NewRequest("POST", "/login", bytes.NewReader([]byte(tt.body)))
			c.Request.Header.Set("Content-Type", "application/json")

			h.Login(c)

			assert.Equal(t, tt.expectedCode, w.Code)
			var resp map[string]string
			json.Unmarshal(w.Body.Bytes(), &resp)
			assert.Contains(t, resp[tt.expectedKey], tt.expectedVal)
		})
	}
}
