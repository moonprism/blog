package cmd

import (
	"fmt"

	"github.com/moonprism/blog/core"
	"github.com/moonprism/blog/model"
	"github.com/urfave/cli/v2"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/ssh/terminal"
)

func NewAdminCommand(app *core.App) *cli.Command {
	command := &cli.Command{
		Name: "admin",
		Subcommands: []*cli.Command{
			{
				Name:  "construct",
				Usage: "sync database struct",
				Action: func(ctx *cli.Context) error {
					return app.O.Sync2(new(model.Article))
				},
			},
			{
				Name:  "passwd",
				Usage: "set account password",
				Action: func(ctx *cli.Context) error {
					fmt.Println("please input password:")
					password, _ := terminal.ReadPassword(0)
					pass, err := bcrypt.GenerateFromPassword(password, 14)
					if err != nil {
						return err
					}
					app.Setting.Account.Pass = string(pass)
					return app.ReSetting()
				},
			},
		},
	}
	return command
}
