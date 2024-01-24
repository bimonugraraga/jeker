package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(plainText string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plainText), 16)
	return string(bytes), err
}

func VerifyPassword(plainText string, hashPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(plainText))
	return err == nil
}
