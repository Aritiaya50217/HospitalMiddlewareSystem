package postgres

import "time"

type PatientModel struct {
	ID           int64  `gorm:"primaryKey"`
	FirstNameTH  string `gorm:"uniqueIndex"`
	MiddleNameTH string
	LastNameTH   string
	FirstNameEN  string
	MiddleNameEN string
	LastNameEN   string
	DataOfBirth  time.Time
	PatientHN    string //  หมายเลขประจำตัวผู้ป่วย
	NationalID   string
	PassportID   string
	PhoneNumber  string
	Email        string
	GenderID     int64
	Gender       GenderModel `gorm:"foreignKey:GenderID"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (PatientModel) TableName() string {
	return "patients"
}
