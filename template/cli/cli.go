// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package cli

import (
	"context"
	"os"

	"github.com/{{github}}/cli/{{toLower child}}"
	"github.com/{{github}}/cli/{{toLower parent}}"
	"github.com/{{github}}/cli/member"
	"github.com/{{github}}/cli/project"
	"github.com/{{github}}/cli/server"
	"github.com/{{github}}/cli/token"
	"github.com/{{github}}/cli/user"
	"github.com/{{github}}/cli/users"
	"github.com/{{github}}/version"

	"gopkg.in/alecthomas/kingpin.v2"
)

// empty context
var nocontext = context.Background()

// application name
var application = "{{app}}"

// application description
var description = "description goes here" // TODO edit this application description

// Command parses the command line arguments and then executes a
// subcommand program.
func Command() {
	app := kingpin.New(application, description)
	server.Register(app)
	user.Register(app)
	project.Register(app)
	{{toLower parent}}.Register(app)
	{{toLower child}}.Register(app)
	member.Register(app)
	users.Register(app)
	token.Register(app)
	registerLogin(app)
	registerLogout(app)
	registerRegister(app)

	kingpin.Version(version.Version.String())
	kingpin.MustParse(app.Parse(os.Args[1:]))
}
