// This file was automatically generated from
//
//	profile.go
//
// by
//
//	generator
//
// at
//
//	2015-07-29T15:00:14+03:00
//
// Do not edit it!

package model

import (
	"errors"

	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

// Save will put this Profile into Datastore using the given key.
func (p_ Profile) Save(ctx context.Context, key ...*datastore.Key) (*datastore.Key, error) {
	if len(key) > 1 {
		panic("zero or one key expected")
	}

	if len(key) == 1 && key[0] != nil {
		return datastore.Put(ctx, key[0], &p_)
	}

	return datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "profiles", nil), &p_)
}

// SaveWithParent can be used to save this Profile as child of another
// entity.
// This will error if parent == nil.
func (p_ Profile) SaveWithParent(ctx context.Context, parent *datastore.Key) (*datastore.Key, error) {
	if parent == nil {
		return nil, errors.New("parent key is nil, expected a valid key")
	}
	return datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "profiles", parent), &p_)
}

// NewQueryForProfile prepares a datastore.Query that can be
// used to query entities of type Profile.
func NewQueryForProfile() *datastore.Query {
	return datastore.NewQuery("profiles")
}
