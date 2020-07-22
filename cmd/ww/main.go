/*
	Wetware - the distributed programming language
	Copyright 2020, Louis Thibault.  All rights reserved.
*/

package main

import (
	"errors"
	"os"

	"github.com/urfave/cli/v2"

	log "github.com/lthibault/log/pkg"

	"github.com/wetware/ww/internal/cmd/boot"
	"github.com/wetware/ww/internal/cmd/client"
	"github.com/wetware/ww/internal/cmd/keygen"
	"github.com/wetware/ww/internal/cmd/start"
)

const version = "0.0.0"

var flags = []cli.Flag{
	&cli.StringFlag{
		Name:    "logfmt",
		Aliases: []string{"f"},
		Usage:   "text, json, none",
		Value:   "text",
		EnvVars: []string{"CASM_LOGFMT"},
	},
	&cli.StringFlag{
		Name:    "loglvl",
		Usage:   "trace, debug, info, warn, error, fatal",
		Value:   "info",
		EnvVars: []string{"CASM_LOGLVL"},
	},
	&cli.BoolFlag{
		Name:    "prettyprint",
		Aliases: []string{"pp"},
		Usage:   "pretty-print JSON output",
		Hidden:  true,
	},
}

var commands = []*cli.Command{
	start.Command(),
	client.Command(),
	keygen.Command(),
	boot.Command(),
}

func main() {
	run(&cli.App{
		Name:                 "wetware",
		Usage:                "the distributed programming language",
		UsageText:            "ww [global options] command [command options] [arguments...]",
		Version:              version,
		EnableBashCompletion: true,
		Flags:                flags,
		Commands:             commands,
		Action:               repl,
	})
}

func repl(c *cli.Context) error {
	return errors.New("shell NOT IMPLEMENTED")
	// return c.App.Command("shell").Run(c)
}

func run(app *cli.App) {
	if err := app.Run(os.Args); err != nil {
		log.New().Fatal(err)
	}
}
