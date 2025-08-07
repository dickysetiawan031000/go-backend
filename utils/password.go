package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	// log hash
	fmt.Println("Hashing password:", password) // Uncomment for debugging

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	// log check
	fmt.Println("Checking password:", password, "against hash:", hash) // Uncomment for debugging
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
