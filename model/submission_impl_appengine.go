// This file was automatically generated from
//
//	submission.go
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

// Save will put this Submission into Datastore using the given key.
func (s_ Submission) Save(ctx context.Context, key ...*datastore.Key) (*datastore.Key, error) {
	if len(key) > 1 {
		panic("zero or one key expected")
	}

	if len(key) == 1 && key[0] != nil {
		return datastore.Put(ctx, key[0], &s_)
	}

	return datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "submissions", nil), &s_)
}

// SaveWithParent can be used to save this Submission as child of another
// entity.
// This will error if parent == nil.
func (s_ Submission) SaveWithParent(ctx context.Context, parent *datastore.Key) (*datastore.Key, error) {
	if parent == nil {
		return nil, errors.New("parent key is nil, expected a valid key")
	}
	return datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "submissions", parent), &s_)
}

// NewQueryForSubmission prepares a datastore.Query that can be
// used to query entities of type Submission.
func NewQueryForSubmission() *datastore.Query {
	return datastore.NewQuery("submissions")
}
