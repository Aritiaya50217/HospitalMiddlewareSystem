package entity

import "time"

type Hospital struct {
	ID        int64     `gorm:"primaryKey;column:id"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
