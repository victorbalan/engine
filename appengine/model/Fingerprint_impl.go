// This file was automatically generated from "../../templates/model/fingerprint.got". Do not edit!

package model

import (
	"encoding/json"
	"net/http"

	"google.golang.org/appengine/datastore"
)

type Fingerprints []Fingerprint

// Write takes a key and the corresponding writes it out to w after marshaling to JSON.
func (f_ Fingerprint) Write(w http.ResponseWriter, key *datastore.Key) {
	body, err := json.Marshal(map[string]Fingerprint{
		key.Encode(): f_,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(body)
}

// Write will write out all Entities to w in JSON format.
func (f_ Fingerprints) Write(w http.ResponseWriter, keys []*datastore.Key) {
	if len(keys) != len(f_) {
		http.Error(w, "length mismatch while writing entities", http.StatusInternalServerError)
		return
	}

	tmp := make(map[string]Fingerprint, len(keys))
	for i := 0; i < len(keys); i++ {
		tmp[keys[i].String()] = f_[i]
	}

	body, err := json.Marshal(tmp)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(body)
}

// NewQueryForFingerprint prepares a datastore.Query that can be
// used to query entities of type Fingerprint.
func NewQueryForFingerprint() *datastore.Query {
	return datastore.NewQuery(FingerprintKind)
}
