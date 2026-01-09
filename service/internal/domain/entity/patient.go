package entity

import "time"

type Patient struct {
	ID           int       `gorm:"primaryKey;column:id"`
	FirstNameTH  string    `gorm:"column:first_name_th"`
	MiddleNameTH string    `gorm:"column:middle_name_th"`
	LastNameTH   string    `gorm:"column:last_name_th"`
	FirstNameEN  string    `gorm:"column:first_name_en"`
	MiddleNameEN string    `gorm:"column:middle_name_en"`
	LastNameEN   string    `gorm:"column:last_name_en"`
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
