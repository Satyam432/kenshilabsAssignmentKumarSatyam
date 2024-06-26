package utils

import (
	"example.com/m/app/models"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateToke(userData models.User) (string, error) {
	// Create a new claims.
	claims := jwt.MapClaims{}

	claims["UserId"] = userData.UserId
	claims["Email"] = userData.Email
	claims["Password"] = userData.Password

	// Create a new JWT access token with claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("secret"))
	// Return the token string.
	if err != nil {
		return "", err
	}
	// Return the token string.
	return t, nil
}
