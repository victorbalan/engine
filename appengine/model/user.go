// This file was automatically generated from "../../templates/model/user.got". Do not edit!

package model

import (
	"encoding/json"
	"errors"
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

const UserKind = "users"

type User struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	HashedPassword []byte `json:"-"`
}

// Save will put this User into Datastore using the given key.
func (x User) Save(ctx context.Context, key ...*datastore.Key) (*datastore.Key, error) {
	if len(key) > 1 {
		panic("zero or one key expected")
	}

	if len(key) == 1 && key[0] != nil {
		return datastore.Put(ctx, key[0], &x)
	}

	return datastore.Put(ctx, datastore.NewIncompleteKey(ctx, UserKind, nil), &x)
}

// SaveWithParent can be used to save this User as child of another
// entity.
// This will error if parent == nil.
func (x User) SaveWithParent(ctx context.Context, parent *datastore.Key) (*datastore.Key, error) {
	if parent == nil {
		return nil, errors.New("parent key is nil, expected a valid key")
	}
	return datastore.Put(ctx, datastore.NewIncompleteKey(ctx, UserKind, parent), &x)
}

// Write takes a key and the corresponding writes it out to w after marshaling to JSON.
func (x User) Write(w http.ResponseWriter, key *datastore.Key) {
	body, err := json.Marshal(map[string]User{
		key.Encode(): x,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(body)
}

type Users []User

// Write will write out all Entities to w in JSON format.
func (x Users) Write(w http.ResponseWriter, keys []*datastore.Key) {
	if len(keys) != len(x) {
		http.Error(w, "length mismatch while writing entities", http.StatusInternalServerError)
		return
	}

	tmp := make(map[string]User, len(keys))
	for i := 0; i < len(keys); i++ {
		tmp[keys[i].String()] = x[i]
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
