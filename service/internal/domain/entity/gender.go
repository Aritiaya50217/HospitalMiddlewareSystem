package entity

import "time"

type Gender struct {
	ID        int       `gorm:"primaryKey;column:id"`
	Name      string    `gorm:"column:name"` // M, F
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
