package postgres

import (
	"strconv"
	"strings"
	"time"

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

func (r *patientRepository) Search(hospitalID int64,nationalID, passportID,firstname, middlename, lastname,dateOfBirth, phoneNumber, email string) ([]entity.Patient, error) {

	var models []entity.Patient
	db := r.db.Model(&PatientModel{}).
		Where("hospital_id = ?", hospitalID)

	// --- Strict identifiers (AND) ---
	if nationalID != "" {
		db = db.Where("national_id = ?", nationalID)
	}
	if passportID != "" {
		db = db.Where("passport_id = ?", passportID)
	}

	// --- Collect flexible OR conditions ---
	var orConditions []string
	var orArgs []interface{}

	if firstname != "" {
		orConditions = append(orConditions, "(first_name_th ILIKE ? OR first_name_en ILIKE ?)")
		orArgs = append(orArgs, "%"+firstname+"%", "%"+firstname+"%")
	}
	if middlename != "" {
		orConditions = append(orConditions, "(middle_name_th ILIKE ? OR middle_name_en ILIKE ?)")
		orArgs = append(orArgs, "%"+middlename+"%", "%"+middlename+"%")
	}
	if lastname != "" {
		orConditions = append(orConditions, "(last_name_th ILIKE ? OR last_name_en ILIKE ?)")
		orArgs = append(orArgs, "%"+lastname+"%", "%"+lastname+"%")
	}
	if phoneNumber != "" {
		orConditions = append(orConditions, "phone_number ILIKE ?")
		orArgs = append(orArgs, "%"+phoneNumber+"%")
	}
	if email != "" {
		orConditions = append(orConditions, "email ILIKE ?")
		orArgs = append(orArgs, "%"+email+"%")
	}
	if dateOfBirth != "" {
		parts := strings.Split(dateOfBirth, "-")
		if len(parts) == 3 {
			yearBE, _ := strconv.Atoi(parts[0])
			month, _ := strconv.Atoi(parts[1])
			day, _ := strconv.Atoi(parts[2])
			yearAD := yearBE - 543

			loc, _ := time.LoadLocation("Asia/Bangkok")
			thaiStart := time.Date(yearAD, time.Month(month), day, 0, 0, 0, 0, loc)
			thaiEnd := thaiStart.Add(24 * time.Hour)

			start := thaiStart.UTC()
			end := thaiEnd.UTC()

			orConditions = append(orConditions, "(date_of_birth >= ? AND date_of_birth < ?)")
			orArgs = append(orArgs, start, end)
		}
	}

	// --- Combine OR conditions if any ---
	if len(orConditions) > 0 {
		combined := strings.Join(orConditions, " OR ")
		db = db.Where(combined, orArgs...)
	}

	// --- Execute query ---
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
