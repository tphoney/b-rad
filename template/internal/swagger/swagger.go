// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

// Package swagger defines the swagger specification.
//
//     Schemes: http, https
//     BasePath: /api/v1
//     Version: 1.0.0
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package swagger

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:generate swagger generate spec -o files/swagger.json

//go:embed files/*
var content embed.FS

// FileSystem provides access to the static web server
// content, embedded in the binary.
func FileSystem() http.FileSystem {
	fsys, _ := fs.Sub(content, "files")
	return http.FS(fsys)
}
