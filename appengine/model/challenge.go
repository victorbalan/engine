// This file was automatically generated from "../../templates/model/challenge.got". Do not edit!

package model

import (
	"encoding/json"
	"errors"
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

// ChallengeKind is the kind used to
// store challenges in Datastore.
const ChallengeKind = "challenges"

// Challenge contains the data of a challenge
// with the company that created it.
type Challenge struct {
	Name         string         `json:"name"`
	Instructions string         `json:"instructions"`
	Company      *datastore.Key `json:"company"`
	WebInterface string         `json:"webInterface"`
	Runner       string         `json:"-"`
	Flags        string         `json:"-"`
	Languages    []string       `json:"languages"`
}

// Save will put this Challenge into Datastore using the given key.
func (x Challenge) Save(ctx context.Context, key ...*datastore.Key) (*datastore.Key, error) {
	if len(key) > 1 {
		panic("zero or one key expected")
	}

	if len(key) == 1 && key[0] != nil {
		return datastore.Put(ctx, key[0], &x)
	}

	return datastore.Put(ctx, datastore.NewIncompleteKey(ctx, ChallengeKind, nil), &x)
}

// SaveWithParent can be used to save this Challenge as child of another
// entity.
// This will error if parent == nil.
func (x Challenge) SaveWithParent(ctx context.Context, parent *datastore.Key) (*datastore.Key, error) {
	if parent == nil {
		return nil, errors.New("parent key is nil, expected a valid key")
	}
	return datastore.Put(ctx, datastore.NewIncompleteKey(ctx, ChallengeKind, parent), &x)
}

// Write takes a key and the corresponding writes it out to w after marshaling to JSON.
func (x Challenge) Write(w http.ResponseWriter, key *datastore.Key) {
	body, err := json.Marshal(map[string]Challenge{
		key.Encode(): x,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(body)
}

type Challenges []Challenge

// Write will write out all Entities to w in JSON format.
func (x Challenges) Write(w http.ResponseWriter, keys []*datastore.Key) {
	if len(keys) != len(x) {
		http.Error(w, "length mismatch while writing entities", http.StatusInternalServerError)
		return
	}

	tmp := make(map[string]Challenge, len(keys))
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

// NewQueryForChallenge prepares a datastore.Query that can be
// used to query entities of type Challenge.
func NewQueryForChallenge() *datastore.Query {
	return datastore.NewQuery(ChallengeKind)
}
