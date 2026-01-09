package entity

import "time"

type StaffPatient struct {
	ID        int64 `gorm:"primaryKey"`
	UserID    int64
	PatientID int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
