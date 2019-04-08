package user

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// User represents a podoloff user
type User struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	LiveToken string
}

// EncryptPassword encrypts password and return as string
func EncryptPassword(password []byte) string {
	hash, _ := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	return string(hash)
}

// ParseUser decodes request body into User struct
func ParseUser(r *http.Request) (*User, error) {
	var u User
	if r.Body == nil {
		return &u, errors.New("No request body")
	}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&u)
	return &u, err
}

// ParseToken finds cookie in request if present
func ParseToken(r *http.Request) (string, error) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		return "", err
	}
	return cookie.Value, err
}

// IssueCookie creates a new cookie
func (u *User) IssueCookie() *http.Cookie {
	token := uuid.Must(uuid.NewV4()).String()
	u.LiveToken = token
	return &http.Cookie{
		Name:    "session_token",
		Value:   token,
		Expires: time.Now().Add(120 * time.Second),
	}
}
