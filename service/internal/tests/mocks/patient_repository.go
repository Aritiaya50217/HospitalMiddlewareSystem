package mocks

import (
	"time"

	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/domain/entity"
)

type PatientRepositoryMock struct {
	SearchFn     func(hospitalID int64, nationalID, passportID, firstname, middlename, lastname, dateOfBirth, phoneNumber, email string) ([]entity.Patient, error)
	SearchByIDFn func(id string, hospitalID int64) (*entity.Patient, error)
}

func (m *PatientRepositoryMock) Search(hospitalID int64, nationalID, passportID, firstname, middlename, lastname, dateOfBirth, phoneNumber, email string) ([]entity.Patient, error) {
	if m.SearchFn != nil {
		return m.SearchFn(hospitalID, nationalID, passportID,
			firstname, middlename, lastname,
			dateOfBirth, phoneNumber, email)
	}
	return []entity.Patient{}, nil
}

func (m *PatientRepositoryMock) SearchByID(id string, hospitalID int64) (*entity.Patient, error) {
	return m.SearchByIDFn(id, hospitalID)
}

func MockPatients() []entity.Patient {
	loc, _ := time.LoadLocation("Asia/Bangkok")
	return []entity.Patient{
		{
			ID:           1,
			FirstNameTH:  "สมชาย",
			MiddleNameTH: "-",
			LastNameTH:   "ใจดี",
			FirstNameEN:  "Somchai",
			MiddleNameEN: "-",
			LastNameEN:   "Jaidee",
			DateOfBirth:  time.Date(1990, 5, 12, 0, 0, 0, 0, loc),
			PatientHN:    "HN0001",
			NationalID:   "1103701234567",
			PassportID:   "123456789",
			PhoneNumber:  "0812345678",
			Email:        "somchai@example.com",
			GenderID:     1,
			HospitalID:   1,
		},
	}
}
