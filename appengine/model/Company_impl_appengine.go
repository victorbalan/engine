// This file was automatically generated from "../../templates/model/company.got". Do not edit!

package model

import (
	"errors"

	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

// Save will put this Company into Datastore using the given key.
func (c_ Company) Save(ctx context.Context, key ...*datastore.Key) (*datastore.Key, error) {
	if len(key) > 1 {
		panic("zero or one key expected")
	}

	if len(key) == 1 && key[0] != nil {
		return datastore.Put(ctx, key[0], &c_)
	}

	return datastore.Put(ctx, datastore.NewIncompleteKey(ctx, CompanyKind, nil), &c_)
}

// SaveWithParent can be used to save this Company as child of another
// entity.
// This will error if parent == nil.
func (c_ Company) SaveWithParent(ctx context.Context, parent *datastore.Key) (*datastore.Key, error) {
	if parent == nil {
		return nil, errors.New("parent key is nil, expected a valid key")
	}
	return datastore.Put(ctx, datastore.NewIncompleteKey(ctx, CompanyKind, parent), &c_)
}
