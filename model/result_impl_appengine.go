// This file was automatically generated from
//
//	result.go
//
// by
//
//	generator
//
// at
//
//	2015-07-29T14:54:30+03:00
//
// Do not edit it!

// +build appenginevm

package model

import (
	"errors"

	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

// Save will put this Result into Datastore using the given key.
func (r_ Result) Save(ctx context.Context, key ...*datastore.Key) (*datastore.Key, error) {
	if len(key) > 1 {
		panic("zero or one key expected")
	}

	if len(key) == 1 && key[0] != nil {
		return datastore.Put(ctx, key[0], &r_)
	}

	return datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "results", nil), &r_)
}

// SaveWithParent can be used to save this Result as child of another
// entity.
// This will error if parent == nil.
func (r_ Result) SaveWithParent(ctx context.Context, parent *datastore.Key) (*datastore.Key, error) {
	if parent == nil {
		return nil, errors.New("parent key is nil, expected a valid key")
	}
	return datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "results", parent), &r_)
}

// NewQueryForResult prepares a datastore.Query that can be
// used to query entities of type Result.
func NewQueryForResult() *datastore.Query {
	return datastore.NewQuery("results")
}
