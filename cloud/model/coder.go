// This file was automatically generated from "../../templates/model/coder.got". Do not edit!

package model

// CoderKind the name of the collection in datastore
const CoderKind = "coders"

// Coder contains the data related to a coder
type Coder struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
