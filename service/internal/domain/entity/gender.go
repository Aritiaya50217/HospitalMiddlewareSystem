package entity

import "time"

type Gender struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`         // Male , Female
	Abbreviation string    `json:"abbreviation"` // M, F
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
