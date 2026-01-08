package entity

import "time"

type Role struct {
	ID        int64     `gorm:"primaryKey;column:id"`
	Name      string    `gorm:"column:name"` // admin, staff
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
