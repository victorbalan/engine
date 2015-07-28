// This file was automatically generated from "../../templates/model/company.got". Do not edit!

package model

import (
	"encoding/json"
	"net/http"

	"google.golang.org/cloud/datastore"
)

type Companys []Company

// Write takes a key and the corresponding writes it out to w after marshaling to JSON.
func (c_ Company) Write(w http.ResponseWriter, key *datastore.Key) {
	body, err := json.Marshal(map[string]Company{
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
func (c_ Companys) Write(w http.ResponseWriter, keys []*datastore.Key) {
	if len(keys) != len(c_) {
		http.Error(w, "length mismatch while writing entities", http.StatusInternalServerError)
		return
	}

	tmp := make(map[string]Company, len(keys))
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

// NewQueryForCompany prepares a datastore.Query that can be
// used to query entities of type Company.
func NewQueryForCompany() *datastore.Query {
	return datastore.NewQuery(CompanyKind)
}
