package postgres

import "time"

type GenderModel struct {
	ID           int64  `gorm:"primaryKey"`
	Name         string `gorm:"uniqueIndex"`         // Male , Female
	Abbreviation string `gorm:"column:abbreviation"` // M, F
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
