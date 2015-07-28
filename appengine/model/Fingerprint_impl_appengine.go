// This file was automatically generated from "../../templates/model/fingerprint.got". Do not edit!

// +build appengine appenginevm

package model

import (
	"errors"

	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

// Save will put this Fingerprint into Datastore using the given key.
func (f_ Fingerprint) Save(ctx context.Context, key ...*datastore.Key) (*datastore.Key, error) {
	if len(key) > 1 {
		panic("zero or one key expected")
	}

	if len(key) == 1 && key[0] != nil {
		return datastore.Put(ctx, key[0], &f_)
	}

	return datastore.Put(ctx, datastore.NewIncompleteKey(ctx, FingerprintKind, nil), &f_)
}

// SaveWithParent can be used to save this Fingerprint as child of another
// entity.
// This will error if parent == nil.
func (f_ Fingerprint) SaveWithParent(ctx context.Context, parent *datastore.Key) (*datastore.Key, error) {
	if parent == nil {
		return nil, errors.New("parent key is nil, expected a valid key")
	}
	return datastore.Put(ctx, datastore.NewIncompleteKey(ctx, FingerprintKind, parent), &f_)
}
