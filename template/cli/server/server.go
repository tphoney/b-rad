// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package server

import (
	"context"
	"os"

	"github.com/{{github}}/internal/logger"
	"github.com/{{github}}/types"
	"github.com/{{github}}/version"

	"github.com/joho/godotenv"
	"github.com/mattn/go-isatty"
	"github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
)

type command struct {
	envfile string
}

func (c *command) run(*kingpin.ParseContext) error {
	// load environment variables from file.
	godotenv.Load(c.envfile)

	// create the system configuration store by loading
	// data from the environment.
	config, err := load()
	if err != nil {
		logrus.Fatal(err)
	}

	// configure the log level
	setupLogger(config)

	server, err := initServer(config)
	if err != nil {
		logrus.Fatal(err)
	}

	// create the http server.
	// server := server.Server{
	// 	Acme:    config.Server.Acme.Enabled,
	// 	Addr:    config.Server.Bind,
	// 	Handler: handler,
	// }

	logrus.
		WithField("revision", version.GitCommit).
		WithField("repository", version.GitRepository).
		WithField("version", version.Version).
		Infof("server listening at %s", config.Server.Bind)

	// starts the http server.
	return server.ListenAndServe(context.Background())
}

// helper function configures the global logger from
// the loaded configuration.
func setupLogger(config *types.Config) {
	logger.L = logrus.NewEntry(
		logrus.StandardLogger(),
	)

	// configure the log level
	switch {
	case config.Trace:
		logrus.SetLevel(logrus.TraceLevel)
	case config.Debug:
		logrus.SetLevel(logrus.DebugLevel)
	}

	// if the terminal is not a tty we should output the
	// logs in json format
	if !isatty.IsTerminal(os.Stdout.Fd()) {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}
}

// Register the server command.
func Register(app *kingpin.Application) {
	c := new(command)

	cmd := app.Command("server", "starts the server").
		Action(c.run)

	cmd.Arg("envfile", "load the environment variable file").
		Default("").
		StringVar(&c.envfile)
}
