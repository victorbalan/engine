// This file was automatically generated from "../../templates/model/challenge.got". Do not edit!

package model

import (
	"errors"

	"golang.org/x/net/context"
	"google.golang.org/cloud/datastore"
)

// Save will put this Challenge into Datastore using the given key.
func (c_ Challenge) Save(c datastore.Client, key ...*datastore.Key) (*datastore.Key, error) {
	ctx := context.Background()
	if len(key) > 1 {
		panic("zero or one key expected")
	}

	if len(key) == 1 && key[0] != nil {
		return c.Put(ctx, key[0], &c_)
	}

	return c.Put(ctx, datastore.NewIncompleteKey(ctx, ChallengeKind, nil), &c_)
}

// SaveWithParent can be used to save this Challenge as child of another
// entity.
// This will error if parent == nil.
func (c_ Challenge) SaveWithParent(c datastore.Client, parent *datastore.Key) (*datastore.Key, error) {
	ctx := context.Background()
	if parent == nil {
		return nil, errors.New("parent key is nil, expected a valid key")
	}
	return c.Put(ctx, datastore.NewIncompleteKey(ctx, ChallengeKind, parent), &c_)
}
