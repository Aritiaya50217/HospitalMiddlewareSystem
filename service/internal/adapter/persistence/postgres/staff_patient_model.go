package postgres

import "time"

type StaffPatient struct {
	ID        int64 `gorm:"primaryKey"`
	UserID    int64
	PatientID int64
	User      UserModel    `gorm:"foreignKey:UserID"`
	Patient   PatientModel `gorm:"foreignKey:PatientID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
