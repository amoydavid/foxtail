package helper

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pwd string) string {

	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {

	}
	return string(hash)
}

func ValidatePasswords(plainPwd string, hashedPwd string) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(plainPwd))
	if err != nil {
		return false
	}
	return true
}
