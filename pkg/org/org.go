package org

import (
	"encoding/json"
	"errors"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Org represents a podoloff organization
type Org struct {
	ID       primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	Name     string               `json:"name"`
	Location string               `json:"location"`
	Website  string               `json:"website"`
	Users    []primitive.ObjectID `json:"users"`
}

// ParseOrg decodes request body into Org struct
func (o *Org) ParseOrg(r *http.Request) error {
	if r.Body == nil {
		return errors.New("No request body")
	}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(o)
	return err
}
