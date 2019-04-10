package utils

import "golang.org/x/crypto/bcrypt"

// EncryptPassword encrypts password and return as string
func EncryptPassword(password []byte) string {
	hash, _ := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	return string(hash)
}
