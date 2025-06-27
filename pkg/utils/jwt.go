package utils

import (
	"time"
	"os"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(adminID uint) (string, error) {
	claims := jwt.MapClaims{
		"admin_id": adminID,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
