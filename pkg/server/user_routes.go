package server

import (
	"context"
	"io"
	"log"
	"net/http"

	"github.com/podoloff/podoloff/pkg/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser handles http request to create a user
func CreateUser(s *Srv, w http.ResponseWriter, r *http.Request) {
	u, err := user.ParseUser(r)
	if err != nil {
		log.Fatal(err)
	}
	u.Password = user.EncryptPassword([]byte(u.Password))

	var result user.User
	err = s.db.Database("podoloff").Collection("users").FindOne(context.TODO(), bson.D{primitive.E{Key: "email", Value: u.Email}}).Decode(&result)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			_, err = s.db.Database("podoloff").Collection("users").InsertOne(context.TODO(), u)
			if err != nil {
				io.WriteString(w, "Unable to add user. Please try again.")
				return
			}
			io.WriteString(w, "User successfully added!")
			return
		}
	}

	io.WriteString(w, "User email already used. Please try again.")
	return
}

// AuthUser authenticates a user
func AuthUser(s *Srv, w http.ResponseWriter, r *http.Request) {
	u, err := user.ParseUser(r)
	if err != nil {
		log.Fatal(err)
	}
	var result user.User
	err = s.db.Database("podoloff").Collection("users").FindOne(context.TODO(), bson.D{primitive.E{Key: "email", Value: u.Email}}).Decode(&result)
	if err != nil {
		io.WriteString(w, "No user found with that email address.")
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(u.Password))

	if err != nil {
		io.WriteString(w, "Passwords do not match.")
		return
	}

	http.SetCookie(w, u.IssueCookie())

	s.cache[u.LiveToken] = u.Email
	io.WriteString(w, "User authenticated.")
	return
}

// AuthTest tests to see if use has active session
func AuthTest(s *Srv, w http.ResponseWriter, r *http.Request) {
	token, _ := user.ParseToken(r)
	io.WriteString(w, "You are authenticated as... "+s.cache[token])
}
