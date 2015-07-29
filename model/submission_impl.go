// This file was automatically generated from
//
//	submission.go
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

type Submissions []Submission

// Write takes a key and the corresponding writes it out to w after marshaling to JSON.
func (s_ Submission) Write(w http.ResponseWriter, key *datastore.Key) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write([]byte(`{"`))
	w.Write([]byte(strconv.FormatInt(key.IntID(), 10)))
	w.Write([]byte(`":`))
	e := json.NewEncoder(w)
	e.Encode(s_)
	w.Write([]byte(`}`))
}

// Write will write out all Entities to w in JSON format.
func (s_ Submissions) Write(w http.ResponseWriter, keys []*datastore.Key) {
	if len(keys) != len(s_) {
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
		e.Encode(s_)
		if i != len(keys) - 1 {
			w.Write([]byte(`,`))
		}
	}
	w.Write([]byte(`}`))
}
