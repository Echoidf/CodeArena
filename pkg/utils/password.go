package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(password []byte) string {
	hashed, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		log.Printf("Error on hashing password, error: %v", err)
		return ""
	}
	return string(hashed)
}
