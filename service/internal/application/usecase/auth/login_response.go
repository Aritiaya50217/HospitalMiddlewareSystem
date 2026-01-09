package login

import "time"

type LoginResponse struct {
	AccessToken string    `json:"access_token"`
	UserID      int64     `json:"user_id"`
	Role        string    `json:"role_name"`
	Hospital    string    `json:"hospital_name"`
	ExpiresAt   time.Time `json:"expires_at"`
}
