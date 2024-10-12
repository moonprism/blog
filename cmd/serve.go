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
			if err := app.InitSetting(); err != nil {
				return err
			}
			if err := app.InitDatabase(); err != nil {
				return err
			}
			if err := app.InitOSS(); err != nil {
				return err
			}
			if err := app.InitCache(); err != nil {
				return err
			}
			if err := app.InitTmpl(); err != nil {
				return err
			}
			app.InitTokenAuth()
			app.InitValidator()
			return api.Serve(app)
		},
	}
	return command
}
