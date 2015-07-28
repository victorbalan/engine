// This file was automatically generated from "../../templates/model/coder.got". Do not edit!

package model

import (
	"encoding/json"
	"net/http"

	"google.golang.org/appengine/datastore"
)

type Coders []Coder

// Write takes a key and the corresponding writes it out to w after marshaling to JSON.
func (c_ Coder) Write(w http.ResponseWriter, key *datastore.Key) {
	body, err := json.Marshal(map[string]Coder{
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
func (c_ Coders) Write(w http.ResponseWriter, keys []*datastore.Key) {
	if len(keys) != len(c_) {
		http.Error(w, "length mismatch while writing entities", http.StatusInternalServerError)
		return
	}

	tmp := make(map[string]Coder, len(keys))
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

// NewQueryForCoder prepares a datastore.Query that can be
// used to query entities of type Coder.
func NewQueryForCoder() *datastore.Query {
	return datastore.NewQuery(CoderKind)
}
