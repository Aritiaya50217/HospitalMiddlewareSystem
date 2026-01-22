package entity

import "time"

const (
	RoleAdmin = "admin"
	RoleStaff = "staff"
)

type Role struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"` // admin, staff
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
