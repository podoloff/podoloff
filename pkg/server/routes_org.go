package server

import (
	"context"
	"io"
	"log"
	"net/http"

	"github.com/podoloff/podoloff/pkg/org"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// create a new organization
func (s *Srv) createOrg(w http.ResponseWriter, r *http.Request) {
	var o org.Org
	err := o.ParseOrg(r)
	if err != nil {
		log.Fatal(err)
	}

	var result org.Org
	err = s.db.Database("podoloff").Collection("orgs").FindOne(context.TODO(), bson.D{primitive.E{Key: "name", Value: o.Name}}).Decode(&result)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			_, err = s.db.Database("podoloff").Collection("orgs").InsertOne(context.TODO(), o)
			if err != nil {
				io.WriteString(w, "Unable to add org. Please try again.")
				return
			}
			io.WriteString(w, "Org successfully added!")
			return
		}
	}

	io.WriteString(w, "An organization with that name already exists. Please try another.")
	return
}

// AuthUser authenticates a user
func (s *Srv) getOrg(w http.ResponseWriter, r *http.Request) {
	var o org.Org
	err := o.ParseOrg(r)
	if err != nil {
		log.Fatal(err)
	}
	var result org.Org
	err = s.db.Database("podoloff").Collection("orgs").FindOne(context.TODO(), bson.D{primitive.E{Key: "name", Value: o.Name}}).Decode(&result)
	if err != nil {
		io.WriteString(w, "No org found with that name.")
		return
	}

	io.WriteString(w, "Found org: "+result.Name)
	return
}
