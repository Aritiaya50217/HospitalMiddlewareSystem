package entity

import "time"

type User struct {
	ID         int64     `json:"id"`
	Username   string    `json:"username"`
	Password   string    `json:"password"` // bcrypt hash
	HospitalID int64     `json:"hospital_id"`
	RoleID     int64     `json:"role_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}