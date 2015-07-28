// This file was automatically generated from "../../templates/model/challenge.got". Do not edit!

package model

import "google.golang.org/cloud/datastore"

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
