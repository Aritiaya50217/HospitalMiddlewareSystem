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

func (r *patientRepository) SearchPatients(id string, hospitalID int64, offset, limit int) ([]*entity.Patient, error) {
	var patients []*entity.Patient

	// db := r.db.Where("hospital_id = ?", hospitalID).Where(`id::text = ? OR national_id = ? OR passport_id = ?`, id, id, id).
	// 	Offset(offset).Limit(limit)

	// if err := db.Find(&patients).Error; err != nil {
	// 	return nil, err
	// }

	return patients, nil
}

func (r *patientRepository) SearchByID(id string, hospitalID int64) (*entity.Patient, error) {
	var patients *entity.Patient

	db := r.db.Where("hospital_id = ?", hospitalID).Where(`id::text = ? OR national_id = ? OR passport_id = ?`, id, id, id)
	if err := db.Find(&patients).Error; err != nil {
		return nil, err
	}

	return patients, nil
}
