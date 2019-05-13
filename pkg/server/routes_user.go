package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/podoloff/podoloff/pkg/user"
	"github.com/podoloff/podoloff/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser handles http request to create a user
func (s *Srv) createUser(w http.ResponseWriter, r *http.Request) {
	var u user.User
	err := u.ParseUser(r)
	if err != nil {
		log.Fatal(err)
	}
	u.Password = utils.EncryptPassword([]byte(u.Password))

	var result user.User
	err = s.db.Database("podoloff").Collection("users").FindOne(context.TODO(), bson.D{primitive.E{Key: "email", Value: u.Email}}).Decode(&result)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			_, err = s.db.Database("podoloff").Collection("users").InsertOne(context.TODO(), u)
			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode("Unable to add user. Please try again.")
				return
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode("User successfully added.")
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("User email already used. Please try again")
	return
}

// AuthUser authenticates a user
func (s *Srv) authUser(w http.ResponseWriter, r *http.Request) {
	var u user.User
	err := u.ParseUser(r)
	if err != nil {
		log.Fatal(err)
	}
	var result user.User
	err = s.db.Database("podoloff").Collection("users").FindOne(context.TODO(), bson.D{primitive.E{Key: "email", Value: u.Email}}).Decode(&result)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("No user found with that email address.")
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(u.Password))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Passwords do not match.")
		return
	}

	http.SetCookie(w, u.IssueCookie())
	log.Print(result.ID)
	s.cache[u.LiveToken] = u.Email
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("User authenticated.")
	return
}

// AuthTest tests to see if use has active session
func (s *Srv) authTest(w http.ResponseWriter, r *http.Request) {
	token, err := utils.ParseCookie(r)
	if err != nil {
		log.Print(err)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Unable to authenticate user.")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("You are authenticated as: " + s.cache[token])
	return
}
