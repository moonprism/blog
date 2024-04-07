package cmd

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/moonprism/blog/core"
	"github.com/moonprism/blog/model"
	"github.com/urfave/cli/v2"
)

func NewDataCommand(app *core.App) *cli.Command {
	command := &cli.Command{
		Name: "data",
		Subcommands: []*cli.Command{
			{
				Name:  "construct",
				Usage: "sync database struct",
				Action: func(ctx *cli.Context) error {
					info, err := app.O.DBMetas()
					spew.Dump(info)
					app.O.Sync2(new(model.Article))
					return err
				},
			},
		},
	}
	return command
}
