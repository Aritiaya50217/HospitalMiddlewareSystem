package postgres

import "time"

type HospitalModel struct {
	ID        int64 `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
