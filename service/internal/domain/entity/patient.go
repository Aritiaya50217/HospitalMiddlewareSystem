package entity

import "time"

type Patient struct {
	ID           int64     `json:"id"`
	FirstNameTH  string    `json:"first_name_th"`
	MiddleNameTH string    `json:"middle_name_th"`
	LastNameTH   string    `json:"last_name_th"`
	FirstNameEN  string    `json:"first_name_en"`
	MiddleNameEN string    `json:"middle_name_en"`
	LastNameEN   string    `json:"last_name_en"`
	DateOfBirth  time.Time `json:"date_of_birth"`
	PatientHN    string    `json:"patient_hn"`
	NationalID   string    `json:"national_id"`
	PassportID   string    `json:"passport_id"`
	PhoneNumber  string    `json:"phone_number"`
	Email        string    `json:"email"`
	GenderID     int       `json:"gender_id"`
	HospitalID   int       `json:"hospital_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
