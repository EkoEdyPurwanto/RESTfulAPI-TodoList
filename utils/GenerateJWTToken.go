package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var secretKey = []byte("my-secret-key")

func GenerateJWTToken(userID int64) (string, error) {
	// Create a new JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign the token with the secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
