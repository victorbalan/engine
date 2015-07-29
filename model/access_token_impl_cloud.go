// This file was automatically generated from
//
//	access_token.go
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

// Save will put this AccessToken into Datastore using the given key.
func (a_ AccessToken) Save(c datastore.Client, key ...*datastore.Key) (*datastore.Key, error) {
	ctx := context.Background()
	if len(key) > 1 {
		panic("zero or one key expected")
	}

	if len(key) == 1 && key[0] != nil {
		return c.Put(ctx, key[0], &a_)
	}

	return c.Put(ctx, datastore.NewIncompleteKey(ctx, "accessTokens", nil), &a_)
}

// SaveWithParent can be used to save this AccessToken as child of another
// entity.
// This will error if parent == nil.
func (a_ AccessToken) SaveWithParent(c datastore.Client, parent *datastore.Key) (*datastore.Key, error) {
	ctx := context.Background()
	if parent == nil {
		return nil, errors.New("parent key is nil, expected a valid key")
	}
	return c.Put(ctx, datastore.NewIncompleteKey(ctx, "accessTokens", parent), &a_)
}

// NewQueryForAccessToken prepares a datastore.Query that can be
// used to query entities of type AccessToken.
func NewQueryForAccessToken() *datastore.Query {
	return datastore.NewQuery("accessTokens")
}
