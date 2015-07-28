// This file was automatically generated from "../../templates/model/template.got". Do not edit!

package model

import "google.golang.org/cloud/datastore"

const TemplateKind = "templates"

// Template contains data about a code template assigned to a challenge
type Template struct {
	Language  string         `json:"language"`
	Path      string         `json:"path"`
	Challenge *datastore.Key `json:"challenge"`
}
