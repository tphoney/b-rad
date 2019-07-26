// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package database

import (
	"context"

	"github.com/{{github}}/internal/store"
	"github.com/{{github}}/internal/store/database/mutex"
	"github.com/{{github}}/types"
)

var _ store.{{parent}}Store = (*{{parent}}StoreSync)(nil)

// New{{parent}}StoreSync returns a new {{parent}}StoreSync.
func New{{parent}}StoreSync(store *{{parent}}Store) *{{parent}}StoreSync {
	return &{{parent}}StoreSync{store}
}

// {{parent}}StoreSync synronizes read and write access to the
// {{toLower parent}} store. This prevents race conditions when the database
// type is sqlite3.
type {{parent}}StoreSync struct{ *{{parent}}Store }

// Find finds the {{toLower parent}} by id.
func (s *{{parent}}StoreSync) Find(ctx context.Context, id int64) (*types.{{parent}}, error) {
	mutex.RLock()
	defer mutex.RUnlock()
	return s.{{parent}}Store.Find(ctx, id)
}

// List returns a list of {{toLower parent}}s.
func (s *{{parent}}StoreSync) List(ctx context.Context, id int64, opts types.Params) ([]*types.{{parent}}, error) {
	mutex.RLock()
	defer mutex.RUnlock()
	return s.{{parent}}Store.List(ctx, id, opts)
}

// Create saves the {{toLower parent}} details.
func (s *{{parent}}StoreSync) Create(ctx context.Context, {{toLower parent}} *types.{{parent}}) error {
	mutex.Lock()
	defer mutex.Unlock()
	return s.{{parent}}Store.Create(ctx, {{toLower parent}})
}

// Update updates the {{toLower parent}} details.
func (s *{{parent}}StoreSync) Update(ctx context.Context, {{toLower parent}} *types.{{parent}}) error {
	mutex.Lock()
	defer mutex.Unlock()
	return s.{{parent}}Store.Update(ctx, {{toLower parent}})
}

// Delete deletes the {{toLower parent}}.
func (s *{{parent}}StoreSync) Delete(ctx context.Context, {{toLower parent}} *types.{{parent}}) error {
	mutex.Lock()
	defer mutex.Unlock()
	return s.{{parent}}Store.Delete(ctx, {{toLower parent}})
}
