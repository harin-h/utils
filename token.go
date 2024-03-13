package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken() (string, error) {
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.SecretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.SecretKey), nil
	})
	if err != nil {
		return false, err
	}
	return token.Valid, nil
}
