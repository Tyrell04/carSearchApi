package hash

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

// A function that take a password and hash it
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if CheckPasswordHash(password, string(bytes)) == false {
		log.Fatal("Hashing failed")
	}
	return string(bytes), err
}

// A funtion that take a hashed password and a password and compare them
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
