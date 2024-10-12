package cmd

import (
	"fmt"
	"github.com/moonprism/blog/core"
	"github.com/moonprism/blog/models"
	"github.com/urfave/cli/v2"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/term"
)

func NewAdminCommand(app *core.App) *cli.Command {
	command := &cli.Command{
		Name: "admin",
		Subcommands: []*cli.Command{
			{
				Name:  "construct",
				Usage: "sync database struct",
				Action: func(ctx *cli.Context) error {
					if err := app.InitSetting(); err != nil {
						return err
					}
					if err := app.InitDatabase(); err != nil {
						return err
					}
					return app.O.OrmClient.AutoMigrate(
						&models.Article{},
						&models.ArticleContent{},
						&models.Tag{},
						&models.ArticleTags{},
						&models.Attachment{},
						&models.Gist{},
						&models.Comment{},
						&models.LoginRecord{},
					)
				},
			},
			{
				Name:  "passwd",
				Usage: "set account password",
				Action: func(ctx *cli.Context) error {
					if err := app.InitSetting(); err != nil {
						return err
					}
					fmt.Println("please input password:")
					password, _ := term.ReadPassword(0)
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
