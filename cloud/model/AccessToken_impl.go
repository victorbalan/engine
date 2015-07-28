// This file was automatically generated from "../../templates/model/access_token.got". Do not edit!

package model

import (
	"encoding/json"
	"net/http"

	"google.golang.org/cloud/datastore"
)

type AccessTokens []AccessToken

// Write takes a key and the corresponding writes it out to w after marshaling to JSON.
func (a_ AccessToken) Write(w http.ResponseWriter, key *datastore.Key) {
	body, err := json.Marshal(map[string]AccessToken{
		key.Encode(): a_,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(body)
}

// Write will write out all Entities to w in JSON format.
func (a_ AccessTokens) Write(w http.ResponseWriter, keys []*datastore.Key) {
	if len(keys) != len(a_) {
		http.Error(w, "length mismatch while writing entities", http.StatusInternalServerError)
		return
	}

	tmp := make(map[string]AccessToken, len(keys))
	for i := 0; i < len(keys); i++ {
		tmp[keys[i].String()] = a_[i]
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
