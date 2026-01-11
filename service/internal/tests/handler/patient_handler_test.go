package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/adapter/http/handler"
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/application/usecase/gender"
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/application/usecase/patient"
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/domain/entity"
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/tests/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSearchPatientByIDHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		hospitalInCtx  bool
		patientResult  *entity.Patient
		patientError   error
		genderResult   *entity.Gender
		genderError    error
		expectedStatus int
		expectedNameTH string
		expectedLastTH string
		expectedGender string
	}{
		{
			name:           "success",
			hospitalInCtx:  true,
			patientResult:  &entity.Patient{FirstNameTH: "สมชาย", LastNameTH: "ใจดี", GenderID: 1},
			genderResult:   &entity.Gender{ID: 1, Abbreviation: "M"},
			expectedStatus: http.StatusOK,
			expectedNameTH: "สมชาย",
			expectedLastTH: "ใจดี",
			expectedGender: "M",
		},
		{
			name:           "unauthorized (no hospital_id)",
			hospitalInCtx:  false,
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:           "patient not found",
			hospitalInCtx:  true,
			patientError:   errors.New("patient not found"),
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "gender lookup error",
			hospitalInCtx:  true,
			patientResult:  &entity.Patient{GenderID: 1},
			genderError:    errors.New("gender not found"),
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			patientRepo := &mocks.PatientRepositoryMock{
				SearchByIDFn: func(id string, hospitalID int64) (*entity.Patient, error) {
					assert.Equal(t, "1", id)
					return &entity.Patient{FirstNameTH: "สมชาย", LastNameTH: "ใจดี", GenderID: 1}, nil
				},
			}

			genderRepo := &mocks.GenderRepositoryMock{
				FindByIDFn: func(id int64) (*entity.Gender, error) {
					assert.Equal(t, int64(1), id)
					return &entity.Gender{
						ID:           1,
						Abbreviation: "M",
					}, nil
				},
			}

			patientUsecase := patient.NewPatientUsecase(patientRepo)
			genderUsecase := gender.NewGenderUsecase(genderRepo)

			h := handler.NewPatientHandler(patientUsecase, genderUsecase)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			req, _ := http.NewRequest("GET", "/patient/search/1", nil)
			c.Request = req
			c.Params = gin.Params{{Key: "id", Value: "1"}}

			// inject middleware value
			c.Set("hospital_id", int64(1))

			h.SearchByID(c)

			assert.Equal(t, http.StatusOK, w.Code)

			var body map[string]any
			json.Unmarshal(w.Body.Bytes(), &body)

			patient := body["patient"].(map[string]interface{})

			assert.Equal(t, "สมชาย", patient["first_name_th"])
			assert.Equal(t, "ใจดี", patient["last_name_th"])
			assert.Equal(t, "M", patient["gender"])

		})
	}

}
