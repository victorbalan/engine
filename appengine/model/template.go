// This file was automatically generated from "../../templates/model/template.got". Do not edit!

package model

import (
	"encoding/json"
	"errors"
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

const TemplateKind = "templates"

// Template contains data about a code template assigned to a challenge
type Template struct {
	Language  string         `json:"language"`
	Path      string         `json:"path"`
	Challenge *datastore.Key `json:"challenge"`
}

// Save will put this Template into Datastore using the given key.
func (x Template) Save(ctx context.Context, key ...*datastore.Key) (*datastore.Key, error) {
	if len(key) > 1 {
		panic("zero or one key expected")
	}

	if len(key) == 1 && key[0] != nil {
		return datastore.Put(ctx, key[0], &x)
	}

	return datastore.Put(ctx, datastore.NewIncompleteKey(ctx, TemplateKind, nil), &x)
}

// SaveWithParent can be used to save this Template as child of another
// entity.
// This will error if parent == nil.
func (x Template) SaveWithParent(ctx context.Context, parent *datastore.Key) (*datastore.Key, error) {
	if parent == nil {
		return nil, errors.New("parent key is nil, expected a valid key")
	}
	return datastore.Put(ctx, datastore.NewIncompleteKey(ctx, TemplateKind, parent), &x)
}

// Write takes a key and the corresponding writes it out to w after marshaling to JSON.
func (x Template) Write(w http.ResponseWriter, key *datastore.Key) {
	body, err := json.Marshal(map[string]Template{
		key.Encode(): x,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(body)
}

type Templates []Template

// Write will write out all Entities to w in JSON format.
func (x Templates) Write(w http.ResponseWriter, keys []*datastore.Key) {
	if len(keys) != len(x) {
		http.Error(w, "length mismatch while writing entities", http.StatusInternalServerError)
		return
	}

	tmp := make(map[string]Template, len(keys))
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

// NewQueryForTemplate prepares a datastore.Query that can be
// used to query entities of type Template.
func NewQueryForTemplate() *datastore.Query {
	return datastore.NewQuery(TemplateKind)
}
