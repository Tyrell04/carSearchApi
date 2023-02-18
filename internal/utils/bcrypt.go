package utils

import (
	"github.com/marcleonschulz/carSearchApi/exception"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword takes a plain password and hashes it
func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	exception.PanicLogging(err)
	return string(bytes)
}

// CheckPasswordHash compares a plain password with a hashed password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
