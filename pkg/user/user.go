package user

import (
	"encoding/json"
	"errors"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// User represents a podoloff user
type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Encrypt password and return as string
func encryptPassword(password []byte) string {
	hash, _ := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	return string(hash)
}

// Decode request body into User struct
func parseUser(r *http.Request) (User, error) {
	var u User
	if r.Body == nil {
		return u, errors.New("No request body")
	}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&u)
	return u, err
}
