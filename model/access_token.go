package model

import "time"

//go:generate generator

// AccessToken encapsulates a string that be used to
// authenticate a User.
type AccessToken struct {
	string
	Scopes       []string
	Description  string
	Creation     time.Time
	Modification time.Time
	Expiry       time.Time
}
