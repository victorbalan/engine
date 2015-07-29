// This file was automatically generated from
//
//	access_token.go
//
// by
//
//	generator
//
// at
//
//	2015-07-29T15:00:14+03:00
//
// Do not edit it!

package model

import (
	"encoding/json"
	"net/http"
	"strconv"

	"google.golang.org/appengine/datastore"
)

type AccessTokens []AccessToken

// Write takes a key and the corresponding writes it out to w after marshaling to JSON.
func (a_ AccessToken) Write(w http.ResponseWriter, key *datastore.Key) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write([]byte(`{"`))
	w.Write([]byte(strconv.FormatInt(key.IntID(), 10)))
	w.Write([]byte(`":`))
	e := json.NewEncoder(w)
	e.Encode(a_)
	w.Write([]byte(`}`))
}

// Write will write out all Entities to w in JSON format.
func (a_ AccessTokens) Write(w http.ResponseWriter, keys []*datastore.Key) {
	if len(keys) != len(a_) {
		http.Error(w, "length mismatch while writing entities", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write([]byte(`{`))
	e := json.NewEncoder(w)
	for i := 0; i < len(keys); i++ {
		w.Write([]byte(`"`))
		w.Write([]byte(strconv.FormatInt(keys[i].IntID(), 10)))
		w.Write([]byte(`":`))
		e.Encode(a_)
		if i != len(keys) - 1 {
			w.Write([]byte(`,`))
		}
	}
	w.Write([]byte(`}`))
}
