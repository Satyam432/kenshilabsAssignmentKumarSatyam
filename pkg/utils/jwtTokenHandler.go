package utils

import (
	"fmt"
	"time"

	"example.com/m/app/models"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateToke(userData models.User) (string, error) {
	//Expiry 1 day
	expirationTime := time.Now().Add(24 * time.Hour)
	// Create JWT claims
	claims := jwt.MapClaims{
		"UserId":   userData.UserId,
		"Email":    userData.Email,
		"Password": userData.Password,
		"exp":      expirationTime.Unix(), // Expiration time in Unix timestamp
	}

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

func InvalidateToken(tokenString string) (string, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil {
		return "", err
	}

	// Extract claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", err
	}

	// Set expiration time to 5 seconds from now
	claims["exp"] = time.Now().Add(5 * time.Second).Unix()

	// Create a new token with updated claims
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string
	tokenString, err = newToken.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ExtractEmailFromToken(tokenString string) (string, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret signing key
		return []byte("secret"), nil
	})

	if err != nil {
		return "", err
	}

	// Extract the claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		email := claims["Email"].(string)
		return email, nil
	} else {
		return "", fmt.Errorf("invalid token")
	}
}
