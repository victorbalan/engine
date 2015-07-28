// This file was automatically generated from "../../templates/model/fingerprint.got". Do not edit!

package model

import "google.golang.org/cloud/datastore"

// FingerprintKind name of the collection in datastore
const FingerprintKind = "fingerprints"

// Fingerprint contains data that links a coder to a challenge
type Fingerprint struct {
	Coder     *datastore.Key `json:"coder"`
	Challenge *datastore.Key `json:"challenge"`
	Token     string         `json:"token"`
}
