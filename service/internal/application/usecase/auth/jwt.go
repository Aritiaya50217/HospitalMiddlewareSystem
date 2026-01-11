package auth

import (
	"time"

	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/domain/entity"
)

type JWTService interface {
	Generate(user *entity.User) (token string, expiresAt time.Time, err error)
	Validate(token string) (*entity.User, error)
}
