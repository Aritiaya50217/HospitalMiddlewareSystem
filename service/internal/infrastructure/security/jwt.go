package security

import (
	"errors"
	"time"

	"github.com/Aritiaya50217/HospitalMiddlewareSystem/internal/domain/entity"
	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	secret string
}

func NewJWTService(secret string) *JWTService {
	return &JWTService{secret: secret}
}

func (j *JWTService) Generate(user *entity.User) (string, time.Time, error) {
	exp := time.Now().Add(time.Hour * 24)

	claims := jwt.MapClaims{
		"user_id":     user.ID,
		"hospital_id": user.HospitalID,
		"exp":         exp.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(j.secret))
	if err != nil {
		return "", time.Time{}, err
	}
	return signed, exp, nil
}

func (j *JWTService) Validate(tokenStr string) (*entity.User, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	user := &entity.User{}

	if id, ok := claims["user_id"].(float64); ok {
		user.ID = int64(id)
	} else {
		return nil, errors.New("user_id not found in token")
	}

	if hospitalID, ok := claims["hospital_id"].(float64); ok {
		user.HospitalID = int64(hospitalID)
	} else {
		return nil, errors.New("hospital_id not found in token")
	}

	return user, nil
}
