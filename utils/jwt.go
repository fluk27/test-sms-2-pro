package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type jwtCustomClaims struct {
	Username string `json:"username"`
	IsActive bool   `json:"isActive"`
	jwt.RegisteredClaims
}

var JwtSecret string

func LoadJwtSecret(jwtKey string) {
	JwtSecret = jwtKey
}
func CreateJWT(id string, username string, isActive bool, expTime time.Duration) (string, error) {
	// Set custom claims
	claims := &jwtCustomClaims{
		Username: username,
		IsActive: isActive,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   id,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expTime)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	accessToken, err := token.SignedString([]byte(JwtSecret))
	if err != nil {
		return "", err
	}
	return accessToken, nil
}
func DecodeJWT(jwtToken string) (jwtCustomClaims, error) {

	var userData jwtCustomClaims
	// 👇
	token, err := jwt.ParseWithClaims(jwtToken, &userData, func(token *jwt.Token) (interface{}, error) {
		return []byte(JwtSecret), nil
	})
	if err != nil {
		return jwtCustomClaims{}, err
	}

	// Checking token validity
	if !token.Valid {
		return jwtCustomClaims{}, errors.New("invalid token")
	}
	return userData, nil
}
