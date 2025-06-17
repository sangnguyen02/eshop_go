package utils

import "golang.org/x/crypto/bcrypt"

func PasswordVerify(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
