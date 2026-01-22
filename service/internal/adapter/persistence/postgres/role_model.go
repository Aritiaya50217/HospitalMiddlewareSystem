package postgres

import "time"

type RoleModel struct {
	ID        int64 `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
