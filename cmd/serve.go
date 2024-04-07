package cmd

import (
	"github.com/moonprism/blog/api"
	"github.com/moonprism/blog/core"
	"github.com/urfave/cli/v2"
)

func NewServeCommand(app *core.App) *cli.Command {
	command := &cli.Command{
		Name:  "serve",
		Usage: "start web serve",
		Action: func(ctx *cli.Context) error {
			api.Serve(app)
			return nil
		},
	}
	return command
}
