package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(config.Password), []byte(password))
	return err == nil
}
