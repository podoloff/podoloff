package user

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents a podoloff user
type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email     string             `json:"email"`
	Password  string             `json:"password"`
	LiveToken string
}

// ParseUser decodes request body into User struct
func (u *User) ParseUser(r *http.Request) error {
	if r.Body == nil {
		return errors.New("No request body")
	}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(u)
	return err
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
