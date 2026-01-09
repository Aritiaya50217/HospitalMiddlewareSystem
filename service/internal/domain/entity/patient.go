package entity

import "time"

type Patient struct {
	ID           int       `gorm:"primaryKey;column:id"`
	FirstNameTh  string    `gorm:"column:first_name_th"`
	MiddleNameTh string    `gorm:"column:middle_name_th"`
	LastNameTh   string    `gorm:"column:last_name_th"`
	FirstNameEn  string    `gorm:"column:first_name_en"`
	MiddleNameEn string    `gorm:"column:middle_name_en"`
	LastNameEn   string    `gorm:"column:last_name_en"`
	DateOfBirth  time.Time `gorm:"column:data_of_birth"`
	PatientHN    string    `gorm:"column:patient_hn"`
	NationalID   string    `gorm:"column:national_id"`
	PassportID   string    `gorm:"column:passport_id"`
	PhoneNumber  string    `gorm:"column:phone_number"`
	Email        string    `gorm:"column:email"`
	GenderID     int       `gorm:"column:gender_id"`
	HospitalID   int       `gorm:"column:hospital_id"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
}
