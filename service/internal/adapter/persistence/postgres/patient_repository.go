package postgres

import (
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/domain/entity"
	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/domain/repository"
	"gorm.io/gorm"
)

type patientRepository struct {
	db *gorm.DB
}

func NewPatientRepository(db *gorm.DB) repository.PatientRepository {
	return &patientRepository{db: db}
}

// func (r *patientRepository) Search(hospitalID int64, nationalID, passportID, firstname, middlename, lastname, dateOfBirth, phoneNumber, email string) ([]entity.Patient, error) {

// 	var models []PatientModel

// 	db := r.db.Model(&PatientModel{}).
// 		Where("hospital_id = ?", hospitalID)

// 	if nationalID != "" {
// 		db = db.Where("national_id = ?", nationalID)
// 	}

// 	if passportID != "" {
// 		db = db.Where("passport_id = ?", passportID)
// 	}

// 	if firstname != "" {
// 		db = db.Where(
// 			"(first_name_th = ? OR first_name_en = ?)",
// 			firstname, firstname,
// 		)
// 	}

// 	if middlename != "" {
// 		db = db.Where(
// 			"(middle_name_th = ? OR middle_name_en = ?)",
// 			middlename, middlename,
// 		)
// 	}

// 	if lastname != "" {
// 		db = db.Where(
// 			"(last_name_th = ? OR last_name_en = ?)",
// 			lastname, lastname,
// 		)
// 	}

// 	if dateOfBirth != "" {
// 		loc, _ := time.LoadLocation("Asia/Bangkok")
// 		t, err := time.ParseInLocation("2006-01-02", dateOfBirth, loc)
// 		if err == nil {
// 			db = db.Where("date_of_birth = ?", t.UTC())
// 		}
// 	}

// 	if phoneNumber != "" {
// 		db = db.Where("phone_number = ?", phoneNumber)
// 	}

// 	if email != "" {
// 		db = db.Where("email = ?", email)
// 	}

// 	if err := db.Find(&models).Error; err != nil {
// 		return nil, err
// 	}

// 	patients := make([]entity.Patient, 0, len(models))
// 	for _, m := range models {
// 		patients = append(patients, entity.Patient{
// 			ID:           int(m.ID),
// 			HospitalID:   int(m.HospitalID),
// 			NationalID:   m.NationalID,
// 			PassportID:   m.PassportID,
// 			FirstNameTH:  m.FirstNameTH,
// 			FirstNameEN:  m.FirstNameEN,
// 			MiddleNameTH: m.MiddleNameTH,
// 			MiddleNameEN: m.MiddleNameEN,
// 			LastNameTH:   m.LastNameTH,
// 			LastNameEN:   m.LastNameEN,
// 			DateOfBirth:  m.DateOfBirth.UTC(),
// 			PhoneNumber:  m.PhoneNumber,
// 			Email:        m.Email,
// 		})
// 	}

// 	return patients, nil
// }

func (r *patientRepository) SearchByID(id string, hospitalID int64) (*entity.Patient, error) {
	var patients *entity.Patient

	db := r.db.Where("hospital_id = ?", hospitalID).Where(`id::text = ? OR national_id = ? OR passport_id = ?`, id, id, id)
	if err := db.First(&patients).Error; err != nil {
		return nil, err
	}

	return patients, nil
}
