package helper

import (
	"crypto/sha256"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pwd string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	return string(hash)
}

func ValidatePasswords(plainPwd string, hashedPwd string) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(plainPwd))
	return err == nil
}

func Sha256(str string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(str)))
}
