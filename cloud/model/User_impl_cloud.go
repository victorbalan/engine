// This file was automatically generated from "../../templates/model/user.got". Do not edit!

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

	return c.Put(ctx, datastore.NewIncompleteKey(ctx, UserKind, nil), &u_)
}

// SaveWithParent can be used to save this User as child of another
// entity.
// This will error if parent == nil.
func (u_ User) SaveWithParent(c datastore.Client, parent *datastore.Key) (*datastore.Key, error) {
	ctx := context.Background()
	if parent == nil {
		return nil, errors.New("parent key is nil, expected a valid key")
	}
	return c.Put(ctx, datastore.NewIncompleteKey(ctx, UserKind, parent), &u_)
}
