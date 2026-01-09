package entity

import "time"

type User struct {
	ID         int64     `gorm:"primaryKey;autoIncrement"`
	Username   string    `gorm:"not null;uniqueIndex"`
	Password   string    `gorm:"not null"` // bcrypt hash
	HospitalID int64     `gorm:"not null;index"`
	RoleID     int64     `gorm:"not null;index"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}