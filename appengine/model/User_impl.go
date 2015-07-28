// This file was automatically generated from "../../templates/model/user.got". Do not edit!

package model

import (
	"encoding/json"
	"net/http"

	"google.golang.org/appengine/datastore"
)

type Users []User

// Write takes a key and the corresponding writes it out to w after marshaling to JSON.
func (u_ User) Write(w http.ResponseWriter, key *datastore.Key) {
	body, err := json.Marshal(map[string]User{
		key.Encode(): u_,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(body)
}

// Write will write out all Entities to w in JSON format.
func (u_ Users) Write(w http.ResponseWriter, keys []*datastore.Key) {
	if len(keys) != len(u_) {
		http.Error(w, "length mismatch while writing entities", http.StatusInternalServerError)
		return
	}

	tmp := make(map[string]User, len(keys))
	for i := 0; i < len(keys); i++ {
		tmp[keys[i].String()] = u_[i]
	}

	body, err := json.Marshal(tmp)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(body)
}

// NewQueryForUser prepares a datastore.Query that can be
// used to query entities of type User.
func NewQueryForUser() *datastore.Query {
	return datastore.NewQuery(UserKind)
}
