// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package database

import (
	"github.com/{{github}}/internal/store"
	"github.com/{{github}}/types"

	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
)

// WireSet provides a wire set for this package
var WireSet = wire.NewSet(
	ProvideDatabase,
	ProvideUserStore,
	ProvideProjectStore,
	ProvideMemberStore,
	Provide{{parent}}Store,
	Provide{{child}}Store,
)

// ProvideDatabase provides a database connection.
func ProvideDatabase(config *types.Config) (*sqlx.DB, error) {
	return Connect(
		config.Database.Driver,
		config.Database.Datasource,
	)
}

// ProvideUserStore provides a user store.
func ProvideUserStore(db *sqlx.DB) store.UserStore {
	switch db.DriverName() {
	case "postgres":
		return NewUserStore(db)
	default:
		return NewUserStoreSync(
			NewUserStore(db),
		)
	}
}

// ProvideProjectStore provides a project store.
func ProvideProjectStore(db *sqlx.DB) store.ProjectStore {
	switch db.DriverName() {
	case "postgres":
		return NewProjectStore(db)
	default:
		return NewProjectStoreSync(
			NewProjectStore(db),
		)
	}
}

// ProvideMemberStore provides a member store.
func ProvideMemberStore(db *sqlx.DB) store.MemberStore {
	switch db.DriverName() {
	case "postgres":
		return NewMemberStore(db)
	default:
		return NewMemberStoreSync(
			NewMemberStore(db),
		)
	}
}

// Provide{{parent}}Store provides a {{toLower parent}} store.
func Provide{{parent}}Store(db *sqlx.DB) store.{{parent}}Store {
	switch db.DriverName() {
	case "postgres":
		return New{{parent}}Store(db)
	default:
		return New{{parent}}StoreSync(
			New{{parent}}Store(db),
		)
	}
}

// Provide{{child}}Store provides a {{toLower child}} store.
func Provide{{child}}Store(db *sqlx.DB) store.{{child}}Store {
	switch db.DriverName() {
	case "postgres":
		return New{{child}}Store(db)
	default:
		return New{{child}}StoreSync(
			New{{child}}Store(db),
		)
	}
}
