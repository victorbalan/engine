// This file was automatically generated from "../../templates/model/fingerprint.got". Do not edit!

package model

import (
	"encoding/json"
	"errors"
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/cloud/datastore"
)

// FingerprintKind name of the collection in datastore
const FingerprintKind = "fingerprints"

// Fingerprint contains data that links a coder to a challenge
type Fingerprint struct {
	Coder     *datastore.Key `json:"coder"`
	Challenge *datastore.Key `json:"challenge"`
	Token     string         `json:"token"`
}

// Save will put this Fingerprint into Datastore using the given key.
func (x Fingerprint) Save(ctx context.Context, key ...*datastore.Key) (*datastore.Key, error) {
	if len(key) > 1 {
		panic("zero or one key expected")
	}

	if len(key) == 1 && key[0] != nil {
		return datastore.Put(ctx, key[0], &x)
	}

	return datastore.Put(ctx, datastore.NewIncompleteKey(ctx, FingerprintKind, nil), &x)
}

// SaveWithParent can be used to save this Fingerprint as child of another
// entity.
// This will error if parent == nil.
func (x Fingerprint) SaveWithParent(ctx context.Context, parent *datastore.Key) (*datastore.Key, error) {
	if parent == nil {
		return nil, errors.New("parent key is nil, expected a valid key")
	}
	return datastore.Put(ctx, datastore.NewIncompleteKey(ctx, FingerprintKind, parent), &x)
}

// Write takes a key and the corresponding writes it out to w after marshaling to JSON.
func (x Fingerprint) Write(w http.ResponseWriter, key *datastore.Key) {
	body, err := json.Marshal(map[string]Fingerprint{
		key.Encode(): x,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(body)
}

type Fingerprints []Fingerprint

// Write will write out all Entities to w in JSON format.
func (x Fingerprints) Write(w http.ResponseWriter, keys []*datastore.Key) {
	if len(keys) != len(x) {
		http.Error(w, "length mismatch while writing entities", http.StatusInternalServerError)
		return
	}

	tmp := make(map[string]Fingerprint, len(keys))
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

// NewQueryForFingerprint prepares a datastore.Query that can be
// used to query entities of type Fingerprint.
func NewQueryForFingerprint() *datastore.Query {
	return datastore.NewQuery(FingerprintKind)
}
