package helper

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
)

func ValidateJWTToken(tokenString string) (*jwt.Token, error) {
	// Parse the JWT token string
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Return the secret key used to sign the token
		return secretKey, nil
	})

	// Check if there was an error parsing the token
	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}
