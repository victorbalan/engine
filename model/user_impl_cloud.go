// This file was automatically generated from
//
//	user.go
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

// +build !appenginevm

package model

import (
	"errors"

	"golang.org/x/net/context"

	"google.golang.org/cloud/datastore"
)

// Save will put this User into Datastore using the given key.
func (u_ User) Save(c datastore.Client, key ...*datastore.Key) (*datastore.Key, error) {
	ctx := context.Background()
	if len(key) > 1 {
		panic("zero or one key expected")
	}

	if len(key) == 1 && key[0] != nil {
		return c.Put(ctx, key[0], &u_)
	}

	return c.Put(ctx, datastore.NewIncompleteKey(ctx, "users", nil), &u_)
}

// SaveWithParent can be used to save this User as child of another
// entity.
// This will error if parent == nil.
func (u_ User) SaveWithParent(c datastore.Client, parent *datastore.Key) (*datastore.Key, error) {
	ctx := context.Background()
	if parent == nil {
		return nil, errors.New("parent key is nil, expected a valid key")
	}
	return c.Put(ctx, datastore.NewIncompleteKey(ctx, "users", parent), &u_)
}

// NewQueryForUser prepares a datastore.Query that can be
// used to query entities of type User.
func NewQueryForUser() *datastore.Query {
	return datastore.NewQuery("users")
}