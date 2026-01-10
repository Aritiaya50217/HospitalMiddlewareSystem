package postgres

import (
	"strconv"
	"strings"

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

func (r *patientRepository) Search(hospitalID int64, nationalID, passportID, firstname, middlename, lastname, dateOfBirth, phoneNumber, email string) ([]entity.Patient, error) {

	var models []entity.Patient
	db := r.db.Model(&PatientModel{}).Where("hospital_id = ?", hospitalID)

	if nationalID != "" {
		db = db.Where("national_id = ?", nationalID)
	}
	if passportID != "" {
		db = db.Where("passport_id = ?", passportID)
	}
	if firstname != "" {
		db = db.Where("(first_name_th = ? OR first_name_en = ?)", firstname, firstname)
	}
	if middlename != "" {
		db = db.Where("(middle_name_th = ? OR middle_name_en = ?)", middlename, middlename)
	}
	if lastname != "" {
		db = db.Where("(last_name_th = ? OR last_name_en = ?)", lastname, lastname)
	}

	if dateOfBirth != "" {
		parts := strings.Split(dateOfBirth, "-") // "2533-09-20"
		if len(parts) == 3 {
			year, _ := strconv.Atoi(parts[0])
			month, _ := strconv.Atoi(parts[1])
			day, _ := strconv.Atoi(parts[2])

			db = db.Where(`
            EXTRACT(YEAR FROM date_of_birth AT TIME ZONE 'Asia/Bangkok') + 543 = ? AND
            EXTRACT(MONTH FROM date_of_birth AT TIME ZONE 'Asia/Bangkok') = ? AND
            EXTRACT(DAY FROM date_of_birth AT TIME ZONE 'Asia/Bangkok') = ?`,
				year, month, day)
		}
	}

	if phoneNumber != "" {
		db = db.Where("phone_number = ?", phoneNumber)
	}
	if email != "" {
		db = db.Where("email = ?", email)
	}

	if err := db.Find(&models).Error; err != nil {
		return nil, err
	}

	return models, nil
}

func (r *patientRepository) SearchByID(id string, hospitalID int64) (*entity.Patient, error) {
	var patients *entity.Patient

	db := r.db.Where("hospital_id = ?", hospitalID).Where(`id::text = ? OR national_id = ? OR passport_id = ?`, id, id, id)
	if err := db.First(&patients).Error; err != nil {
		return nil, err
	}

	return patients, nil
}
