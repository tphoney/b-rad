// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package memory

import (
	"github.com/{{github}}/internal/store"
	"github.com/google/wire"
)

// WireSet provides a wire set for this package
var WireSet = wire.NewSet(
	New,
	wire.Bind(new(store.SystemStore), new(*SystemStore)),
)
