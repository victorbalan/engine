// +build !appenginevm

package model

import "google.golang.org/cloud/datastore"

// Key is a key suitable for use
// with Google Datastore.
type Key struct {
	*datastore.Key
}
