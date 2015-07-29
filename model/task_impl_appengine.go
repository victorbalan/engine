// This file was automatically generated from
//
//	task.go
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

// Save will put this Task into Datastore using the given key.
func (t_ Task) Save(ctx context.Context, key ...*datastore.Key) (*datastore.Key, error) {
	if len(key) > 1 {
		panic("zero or one key expected")
	}

	if len(key) == 1 && key[0] != nil {
		return datastore.Put(ctx, key[0], &t_)
	}

	return datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "tasks", nil), &t_)
}

// SaveWithParent can be used to save this Task as child of another
// entity.
// This will error if parent == nil.
func (t_ Task) SaveWithParent(ctx context.Context, parent *datastore.Key) (*datastore.Key, error) {
	if parent == nil {
		return nil, errors.New("parent key is nil, expected a valid key")
	}
	return datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "tasks", parent), &t_)
}

// NewQueryForTask prepares a datastore.Query that can be
// used to query entities of type Task.
func NewQueryForTask() *datastore.Query {
	return datastore.NewQuery("tasks")
}
