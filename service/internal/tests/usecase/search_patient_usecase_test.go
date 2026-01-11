package usecase

import (
	"strconv"
	"strings"
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

func TestSearchPatientUsecase(t *testing.T) {
	mockPatients := mocks.MockPatients()

	repo := &mocks.PatientRepositoryMock{
		SearchFn: func(hospitalID int64, nationalID, passportID, firstname, middlename, lastname, dateOfBirth, phoneNumber, email string) ([]entity.Patient, error) {
			var result []entity.Patient
			for _, patient := range mockPatients {
				if hospitalID != int64(patient.HospitalID) {
					continue
				}

				match := true
				if nationalID != "" && nationalID != patient.NationalID {
					match = false
				}

				if passportID != "" && passportID != patient.PassportID {
					match = false
				}

				if firstname != "" && firstname != patient.FirstNameTH && firstname != patient.FirstNameEN {
					match = false
				}

				if middlename != "" && middlename != patient.MiddleNameTH && middlename != patient.MiddleNameEN {
					match = false
				}

				if lastname != "" && lastname != patient.LastNameTH && lastname != patient.LastNameEN {
					match = false
				}

				if phoneNumber != "" && phoneNumber != patient.PhoneNumber {
					match = false
				}

				if email != "" && email != patient.Email {
					match = false
				}

				if dateOfBirth != "" {
					parts := strings.Split(dateOfBirth, "-")
					if len(parts) == 3 {
						year, _ := strconv.Atoi(parts[0])
						month, _ := strconv.Atoi(parts[1])
						day, _ := strconv.Atoi(parts[2])

						yearAD := year - 543
						loc, _ := time.LoadLocation("Asia/Bangkok")
						start := time.Date(yearAD, time.Month(month), day, 0, 0, 0, 0, loc)
						end := start.Add(24 * time.Hour)

						if !(patient.DateOfBirth.Equal(start) || (patient.DateOfBirth.After(start) && patient.DateOfBirth.Before(end))) {
							match = false
						}
					} else {
						match = false
					}
				}

				if match {
					result = append(result, patient)
				}
			}
			return result, nil
		},
	}

	uc := patient.NewPatientUsecase(repo)

	tests := []struct {
		name                string
		request             patient.SearchPatientRequest
		expectedCount       int
		expectedFirstNameTH []string
		expectedLastNameTH  []string
	}{
		{
			name: "search by firstname",
			request: patient.SearchPatientRequest{
				FirstName: "Somchai",
			},
			expectedCount:       1,
			expectedFirstNameTH: []string{"สมชาย"},
			expectedLastNameTH:  []string{"ใจดี"},
		},
		{
			name: "search by lastname",
			request: patient.SearchPatientRequest{
				LastName: "Jaidee",
			},
			expectedCount:       1,
			expectedFirstNameTH: []string{"สมชาย"},
			expectedLastNameTH:  []string{"ใจดี"},
		},
		{
			name: "search by nationalID",
			request: patient.SearchPatientRequest{
				NationalID: "1103701234567",
			},
			expectedCount:       1,
			expectedFirstNameTH: []string{"สมชาย"},
			expectedLastNameTH:  []string{"ใจดี"},
		},
		{
			name: "search by passportID",
			request: patient.SearchPatientRequest{
				PassportID: "123456789",
			},
			expectedCount:       1,
			expectedFirstNameTH: []string{"สมชาย"},
			expectedLastNameTH:  []string{"ใจดี"},
		},
		{
			name: "search by email",
			request: patient.SearchPatientRequest{
				Email: "somchai@example.com",
			},
			expectedCount:       1,
			expectedFirstNameTH: []string{"สมชาย"},
			expectedLastNameTH:  []string{"ใจดี"},
		},
		{
			name: "search by phoneNumber",
			request: patient.SearchPatientRequest{
				PhoneNumber: "0812345678",
			},
			expectedCount:       1,
			expectedFirstNameTH: []string{"สมชาย"},
			expectedLastNameTH:  []string{"ใจดี"},
		},
		{
			name: "search by dateOfBirth",
			request: patient.SearchPatientRequest{
				DateOfBirth: "2533-05-12", // BE
			},
			expectedCount:       1,
			expectedFirstNameTH: []string{"สมชาย"},
			expectedLastNameTH:  []string{"ใจดี"},
		},
		{
			name: "search by multiple fields",
			request: patient.SearchPatientRequest{
				FirstName:   "Somchai",
				LastName:    "ใจดี",
				Email:       "somchai@example.com",
				DateOfBirth: "2533-05-12",
			},
			expectedCount:       1,
			expectedFirstNameTH: []string{"สมชาย"},
			expectedLastNameTH:  []string{"ใจดี"},
		},
		{
			name: "no match",
			request: patient.SearchPatientRequest{
				FirstName: "NonExist",
			},
			expectedCount:       0,
			expectedFirstNameTH: []string{"สมชาย"},
			expectedLastNameTH:  []string{"ใจดี"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := uc.Search(
				1,
				tt.request.NationalID,
				tt.request.PassportID,
				tt.request.FirstName,
				tt.request.MiddleName,
				tt.request.LastName,
				tt.request.DateOfBirth,
				tt.request.PhoneNumber,
				tt.request.Email,
			)

			assert.NoError(t, err)
			assert.Len(t, result, tt.expectedCount)

			for i, p := range result {
				assert.Equal(t, tt.expectedFirstNameTH[i], p.FirstNameTH)
				assert.Equal(t, tt.expectedLastNameTH[i], p.LastNameTH)
			}
		})
	}

}
