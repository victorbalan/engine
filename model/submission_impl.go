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
//	2015-07-29T12:56:49+03:00
//
// Do not edit it!

package model

import (
	"encoding/json"
	"net/http"
)

type Submissions []Submission

// Write takes a key and the corresponding writes it out to w after marshaling to JSON.
func (s_ Submission) Write(w http.ResponseWriter, key Key) {
	body, err := json.Marshal(map[string]Submission{
		key.Encode(): s_,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(body)
}

// Write will write out all Entities to w in JSON format.
func (s_ Submissions) Write(w http.ResponseWriter, keys []Key) {
	if len(keys) != len(s_) {
		http.Error(w, "length mismatch while writing entities", http.StatusInternalServerError)
		return
	}

	tmp := make(map[string]Submission, len(keys))
	for i := 0; i < len(keys); i++ {
		tmp[keys[i].String()] = s_[i]
	}

	body, err := json.Marshal(tmp)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(body)
}
