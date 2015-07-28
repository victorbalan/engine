// This file was automatically generated from "../../templates/model/challenge.got". Do not edit!

package model

import (
	"encoding/json"
	"net/http"

	"google.golang.org/appengine/datastore"
)

type Challenges []Challenge

// Write takes a key and the corresponding writes it out to w after marshaling to JSON.
func (c_ Challenge) Write(w http.ResponseWriter, key *datastore.Key) {
	body, err := json.Marshal(map[string]Challenge{
		key.Encode(): c_,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(body)
}

// Write will write out all Entities to w in JSON format.
func (c_ Challenges) Write(w http.ResponseWriter, keys []*datastore.Key) {
	if len(keys) != len(c_) {
		http.Error(w, "length mismatch while writing entities", http.StatusInternalServerError)
		return
	}

	tmp := make(map[string]Challenge, len(keys))
	for i := 0; i < len(keys); i++ {
		tmp[keys[i].String()] = c_[i]
	}

	body, err := json.Marshal(tmp)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(body)
}

// NewQueryForChallenge prepares a datastore.Query that can be
// used to query entities of type Challenge.
func NewQueryForChallenge() *datastore.Query {
	return datastore.NewQuery(ChallengeKind)
}
