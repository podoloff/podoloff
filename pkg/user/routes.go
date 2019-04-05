package user

import (
	"context"
	"io"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateUser handles http request to create a user
func CreateUser(db *mongo.Client, w http.ResponseWriter, r *http.Request) {
	user, err := parseUser(r)
	if err != nil {
		log.Fatal(err)
	}
	user.Password = encryptPassword([]byte(user.Password))

	var result User
	err = db.Database("podoloff").Collection("users").FindOne(context.TODO(), bson.D{primitive.E{Key: "email", Value: user.Email}}).Decode(&result)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			_, err = db.Database("podoloff").Collection("users").InsertOne(context.TODO(), user)
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
func AuthUser(db *mongo.Client, w http.ResponseWriter, r *http.Request) {
	user, err := parseUser(r)
	if err != nil {
		log.Fatal(err)
	}
	var result User
	err = db.Database("podoloff").Collection("users").FindOne(context.TODO(), bson.D{primitive.E{Key: "email", Value: user.Email}}).Decode(&result)
	if err != nil {
		io.WriteString(w, "No user found with that email address.")
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password))

	if err != nil {
		io.WriteString(w, "Passwords do not match.")
		return
	}
	io.WriteString(w, "User authenticated.")
	return
}
