// This file was automatically generated from "../../templates/model/coder.got". Do not edit!

package model

import (
	"encoding/json"
	"errors"
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

// CoderKind the name of the collection in datastore
const CoderKind = "coders"

// Coder contains the data related to a coder
type Coder struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// Save will put this Coder into Datastore using the given key.
func (x Coder) Save(ctx context.Context, key ...*datastore.Key) (*datastore.Key, error) {
	if len(key) > 1 {
		panic("zero or one key expected")
	}

	if len(key) == 1 && key[0] != nil {
		return datastore.Put(ctx, key[0], &x)
	}

	return datastore.Put(ctx, datastore.NewIncompleteKey(ctx, CoderKind, nil), &x)
}

// SaveWithParent can be used to save this Coder as child of another
// entity.
// This will error if parent == nil.
func (x Coder) SaveWithParent(ctx context.Context, parent *datastore.Key) (*datastore.Key, error) {
	if parent == nil {
		return nil, errors.New("parent key is nil, expected a valid key")
	}
	return datastore.Put(ctx, datastore.NewIncompleteKey(ctx, CoderKind, parent), &x)
}

// Write takes a key and the corresponding writes it out to w after marshaling to JSON.
func (x Coder) Write(w http.ResponseWriter, key *datastore.Key) {
	body, err := json.Marshal(map[string]Coder{
		key.Encode(): x,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(body)
}

type Coders []Coder

// Write will write out all Entities to w in JSON format.
func (x Coders) Write(w http.ResponseWriter, keys []*datastore.Key) {
	if len(keys) != len(x) {
		http.Error(w, "length mismatch while writing entities", http.StatusInternalServerError)
		return
	}

	tmp := make(map[string]Coder, len(keys))
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

// NewQueryForCoder prepares a datastore.Query that can be
// used to query entities of type Coder.
func NewQueryForCoder() *datastore.Query {
	return datastore.NewQuery(CoderKind)
}
