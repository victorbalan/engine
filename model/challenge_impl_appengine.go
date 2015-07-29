// This file was automatically generated from
//
//	challenge.go
//
// by
//
//	generator
//
// at
//
//	2015-07-29T12:56:49+03:00
//
// Do not edit it!

// +build appenginevm

package model

import (
	"errors"

	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

// Save will put this Challenge into Datastore using the given key.
func (c_ Challenge) Save(ctx context.Context, key ...*datastore.Key) (*datastore.Key, error) {
	if len(key) > 1 {
		panic("zero or one key expected")
	}

	if len(key) == 1 && key[0] != nil {
		return datastore.Put(ctx, key[0], &c_)
	}

	return datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "challenges", nil), &c_)
}

// SaveWithParent can be used to save this Challenge as child of another
// entity.
// This will error if parent == nil.
func (c_ Challenge) SaveWithParent(ctx context.Context, parent *datastore.Key) (*datastore.Key, error) {
	if parent == nil {
		return nil, errors.New("parent key is nil, expected a valid key")
	}
	return datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "challenges", parent), &c_)
}

// NewQueryForChallenge prepares a datastore.Query that can be
// used to query entities of type Challenge.
func NewQueryForChallenge() *datastore.Query {
	return datastore.NewQuery("challenges")
}
