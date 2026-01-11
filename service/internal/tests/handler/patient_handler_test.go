package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

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

func TestSearchPatientHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Mock patients
	mockPatients := mocks.MockPatients()

	// Patient repository mock
	patientRepo := &mocks.PatientRepositoryMock{
		SearchFn: func(hospitalID int64, nationalID, passportID, firstname, middlename, lastname, dateOfBirth, phoneNumber, email string) ([]entity.Patient, error) {
			var result []entity.Patient
			for _, p := range mockPatients {
				if hospitalID != int64(p.HospitalID) {
					continue
				}
				match := true
				if nationalID != "" && nationalID != p.NationalID {
					match = false
				}
				if passportID != "" && passportID != p.PassportID {
					match = false
				}
				if firstname != "" && firstname != p.FirstNameTH && firstname != p.FirstNameEN {
					match = false
				}
				if middlename != "" && middlename != p.MiddleNameTH && middlename != p.MiddleNameEN {
					match = false
				}
				if lastname != "" && lastname != p.LastNameTH && lastname != p.LastNameEN {
					match = false
				}
				if phoneNumber != "" && phoneNumber != p.PhoneNumber {
					match = false
				}
				if email != "" && email != p.Email {
					match = false
				}
				if dateOfBirth != "" {
					// Assume dateOfBirth is BE yyyy-mm-dd
					year := p.DateOfBirth.Year() + 543
					month := int(p.DateOfBirth.Month())
					day := p.DateOfBirth.Day()
					expected := fmt.Sprintf("%04d-%02d-%02d", year, month, day)
					if expected != dateOfBirth {
						match = false
					}
				}
				if match {
					result = append(result, p)
				}
			}
			return result, nil
		},
		SearchByIDFn: func(id string, hospitalID int64) (*entity.Patient, error) {
			if id == "1" && hospitalID == 1 {
				return &entity.Patient{
					FirstNameTH: "สมชาย", LastNameTH: "ใจดี", GenderID: 1,
					PatientHN: "HN001", NationalID: "1103701234567", PassportID: "123456789",
					PhoneNumber: "0812345678", Email: "somchai@example.com",
					DateOfBirth: time.Date(1990, 5, 12, 0, 0, 0, 0, time.UTC),
				}, nil
			}
			return nil, errors.New("not found")
		},
	}

	// Gender repository mock
	genderRepo := &mocks.GenderRepositoryMock{
		FindByIDFn: func(id int64) (*entity.Gender, error) {
			if id == 1 {
				return &entity.Gender{ID: 1, Abbreviation: "M"}, nil
			}
			return nil, errors.New("gender not found")
		},
	}

	// Usecases
	ucPatient := patient.NewPatientUsecase(patientRepo)
	ucGender := gender.NewGenderUsecase(genderRepo)

	h := handler.NewPatientHandler(ucPatient, ucGender)

	t.Run("Search Handler", func(t *testing.T) {
		tests := []struct {
			name          string
			query         string
			hospitalID    int64
			expectedCode  int
			expectedCount int
		}{
			{"search by firstname", "first_name=Somchai", 1, http.StatusOK, 1},
			{"search by middlename", "middle_name=-", 1, http.StatusOK, 1},
			{"search by lastname", "last_name=ใจดี", 1, http.StatusOK, 1},
			{"search by nationalID", "national_id=1103701234567", 1, http.StatusOK, 1},
			{"search by passportID", "passport_id=123456789", 1, http.StatusOK, 1},
			{"search by email", "email=somchai@example.com", 1, http.StatusOK, 1},
			{"search by phone", "phone_number=0812345678", 1, http.StatusOK, 1},
			{"search by dob", "date_of_birth=2533-05-12", 1, http.StatusOK, 1},
			{"search by multiple fields", "first_name=Somchai&middle_name=-&last_name=ใจดี&email=somchai@example.com&date_of_birth=2533-05-12", 1, http.StatusOK, 1},
			{"no match", "first_name=NonExist", 1, http.StatusOK, 0},
			{"unauthorized", "first_name=Somchai", 0, http.StatusUnauthorized, 0},
			{"empty query", "", 1, http.StatusBadRequest, 0},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request, _ = http.NewRequest("GET", "/patient/search?"+tt.query, nil)
				if tt.hospitalID != 0 {
					c.Set("hospital_id", tt.hospitalID)
				}

				h.Search(c)

				assert.Equal(t, tt.expectedCode, w.Code)

				if w.Code == http.StatusOK {
					var body map[string]interface{}
					json.Unmarshal(w.Body.Bytes(), &body)
					patients, _ := body["patients"].([]interface{})
					assert.Len(t, patients, tt.expectedCount)

				} else if w.Code == http.StatusBadRequest {
					var body map[string]interface{}
					json.Unmarshal(w.Body.Bytes(), &body)
					assert.Contains(t, body["error"], "required")
					
				} else if w.Code == http.StatusUnauthorized {
					var body map[string]interface{}
					json.Unmarshal(w.Body.Bytes(), &body)
					assert.Contains(t, body["error"], "unauthorized")
				}
			})
		}
	})
}

