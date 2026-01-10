package usecase

import (
	"strconv"
	"testing"
	"time"

	patient "github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/application/usecase/patient"
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/domain/entity"
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestSearchPatientByIDUsecase(t *testing.T) {
	dateOfBirth, _ := time.Parse(time.RFC3339, "1990-05-11T17:00:00Z")

	tests := []struct {
		name          string
		request       patient.SearchPatientRequest
		mockResult    *entity.Patient
		mockError     error
		expectError   bool
		expectedCount int
	}{
		{
			name: "search by national_id",
			request: patient.SearchPatientRequest{
				NationalID: "1103701234567",
			},
			mockResult: &entity.Patient{
				FirstNameTH:  "สมชาย",
				MiddleNameTH: "-",
				LastNameTH:   "ใจดี",
				FirstNameEN:  "Somchai",
				MiddleNameEN: "-",
				LastNameEN:   "Jaidee",
				DateOfBirth:  dateOfBirth,
				PatientHN:    "HN0001",
				NationalID:   "1103707654321",
				PassportID:   "123456789",
				PhoneNumber:  "0812345678",
				Email:        "somchai@example.com",
				GenderID:     1,
			},
		},
		{
			name: "search by passport_id",
			request: patient.SearchPatientRequest{
				PassportID: "P123456",
			},
			mockResult: &entity.Patient{
				FirstNameTH:  "สมชาย",
				MiddleNameTH: "-",
				LastNameTH:   "ใจดี",
				FirstNameEN:  "Somchai",
				MiddleNameEN: "-",
				LastNameEN:   "Jaidee",
				DateOfBirth:  dateOfBirth,
				PatientHN:    "HN0001",
				NationalID:   "1103707654321",
				PassportID:   "123456789",
				PhoneNumber:  "0812345678",
				Email:        "somchai@example.com",
				GenderID:     1,
			},
		},
		{
			name: "search by patient id",
			request: patient.SearchPatientRequest{
				PatientID: 1,
			},
			mockResult: &entity.Patient{
				FirstNameTH:  "สมชาย",
				MiddleNameTH: "-",
				LastNameTH:   "ใจดี",
				FirstNameEN:  "Somchai",
				MiddleNameEN: "-",
				LastNameEN:   "Jaidee",
				DateOfBirth:  dateOfBirth,
				PatientHN:    "HN0001",
				NationalID:   "1103707654321",
				PassportID:   "123456789",
				PhoneNumber:  "0812345678",
				Email:        "somchai@example.com",
				GenderID:     1,
			},
		},
		{
			name: "patient not found",
			request: patient.SearchPatientRequest{
				PatientID: 1,
			},
			mockResult:  nil,
			mockError:   patient.ErrPatientNotFound,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// mock repo
			repo := &mocks.PatientRepositoryMock{
				SearchByIDFn: func(id string, hospitalID int64) (*entity.Patient, error) {
					if id == "0" { // case patient not found
						return nil, patient.ErrPatientNotFound
					}

					if tt.request.PatientID != 0 {
						assert.Equal(t, strconv.FormatInt(tt.request.PatientID, 10), id)
					} else if tt.request.NationalID != "" {
						assert.Equal(t, tt.request.NationalID, id)
					} else if tt.request.PassportID != "" {
						assert.Equal(t, tt.request.PassportID, id)
					}

					return tt.mockResult, tt.mockError
				},
			}

			uc := patient.NewPatientUsecase(repo)

			// act
			var searchID string
			if tt.request.PatientID != 0 {
				searchID = strconv.FormatInt(tt.request.PatientID, 10)
			} else if tt.request.NationalID != "" {
				searchID = tt.request.NationalID
			} else if tt.request.PassportID != "" {
				searchID = tt.request.PassportID
			}

			result, err := uc.SearchByID(searchID, 1)

			// assert
			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, tt.mockResult, result)
			}
		})
	}
}
