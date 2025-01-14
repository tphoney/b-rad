// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

//+build wireinject

package server

import (
	"github.com/{{github}}/internal/router"
	"github.com/{{github}}/internal/server"
	"github.com/{{github}}/internal/store/database"
	"github.com/{{github}}/internal/store/memory"
	"github.com/{{github}}/types"

	"github.com/google/wire"
)

func initServer(config *types.Config) (*server.Server, error) {
	wire.Build(
		database.WireSet,
		memory.WireSet,
		router.WireSet,
		server.WireSet,
	)
	return &server.Server{}, nil
}
