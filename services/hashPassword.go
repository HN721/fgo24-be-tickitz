package services

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash string: %v", err)
	}
	return string(hashed), nil
}
func ComparePassword(password, plain string) error {
	return bcrypt.CompareHashAndPassword([]byte(password), []byte(plain))
}
