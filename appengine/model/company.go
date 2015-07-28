// This file was automatically generated from "../../templates/model/company.got". Do not edit!

package model

// CompanyKind is the kind used to store companies in
// Datastore.
const CompanyKind = "companies"

// Company contains the data related to a company.
type Company struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	HashedPassword []byte `json:"-"`
}
