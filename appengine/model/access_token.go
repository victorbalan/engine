// This file was automatically generated from "../../templates/model/access_token.got". Do not edit!

package model

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

const AccessTokenKind = "accesstokens"

type AccessToken struct {
	Value        []byte    `json:"value"`
	Scopes       []string  `json:"scopes"`
	Description  string    `json:"description"`
	Creation     time.Time `json:"creation"`
	Modification time.Time `json:"modification"`
	Expiry       time.Time `json:"expiry"`
}

// Save will put this AccessToken into Datastore using the given key.
func (x AccessToken) Save(ctx context.Context, key ...*datastore.Key) (*datastore.Key, error) {
	if len(key) > 1 {
		panic("zero or one key expected")
	}

	if len(key) == 1 && key[0] != nil {
		return datastore.Put(ctx, key[0], &x)
	}

	return datastore.Put(ctx, datastore.NewIncompleteKey(ctx, AccessTokenKind, nil), &x)
}

// SaveWithParent can be used to save this AccessToken as child of another
// entity.
// This will error if parent == nil.
func (x AccessToken) SaveWithParent(ctx context.Context, parent *datastore.Key) (*datastore.Key, error) {
	if parent == nil {
		return nil, errors.New("parent key is nil, expected a valid key")
	}
	return datastore.Put(ctx, datastore.NewIncompleteKey(ctx, AccessTokenKind, parent), &x)
}

// Write takes a key and the corresponding writes it out to w after marshaling to JSON.
func (x AccessToken) Write(w http.ResponseWriter, key *datastore.Key) {
	body, err := json.Marshal(map[string]AccessToken{
		key.Encode(): x,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(body)
}

type AccessTokens []AccessToken

// Write will write out all Entities to w in JSON format.
func (x AccessTokens) Write(w http.ResponseWriter, keys []*datastore.Key) {
	if len(keys) != len(x) {
		http.Error(w, "length mismatch while writing entities", http.StatusInternalServerError)
		return
	}

	tmp := make(map[string]AccessToken, len(keys))
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

// NewQueryForAccessToken prepares a datastore.Query that can be
// used to query entities of type AccessToken.
func NewQueryForAccessToken() *datastore.Query {
	return datastore.NewQuery(AccessTokenKind)
}
