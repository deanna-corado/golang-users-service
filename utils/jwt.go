package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(clientID string) (string, error) {
	expDuration, err := time.ParseDuration(os.Getenv("TOKEN_EXP"))
	if err != nil {
		expDuration = 24 * time.Hour // fallback
	}

	claims := jwt.MapClaims{
		"client_id": clientID,
		"exp":       time.Now().Add(expDuration).Unix(),
		"iat":       time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
