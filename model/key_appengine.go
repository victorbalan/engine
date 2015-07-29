// +build appenginevm

package model

import "google.golang.org/appengine/datastore"

// Key is a key suitable for use
// with Google Datastore on App Engine.
type Key struct {
	*datastore.Key
}
