package postgres

import "time"

type UserModel struct {
	ID         int64  `gorm:"primaryKey"`
	Username   string `gorm:"uniqueIndex"`
	Password   string
	HospitalID int64
	RoleID     int64
	Role       RoleModel     `gorm:"foreignKey:RoleID"`
	Hospital   HospitalModel `gorm:"foreignKey:HospitalID"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (UserModel) TableName() string {
	return "users"
}
