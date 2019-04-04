package user

import (
	"context"
	"io"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

// CreateUser handles http request to create a user
func CreateUser(db *mongo.Client, w http.ResponseWriter, r *http.Request) {
	user, err := parseUser(r)
	if err != nil {
		log.Fatal(err)
	}
	user.Password = encryptPassword([]byte(user.Password))
	_, err = db.Database("charon").Collection("users").InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}
	io.WriteString(w, "User successfully added.")
}

// AuthUser authenticates a user
func AuthUser(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "User created.")
}
