package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

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

					err := app.O.OrmClient.AutoMigrate(
						&models.Article{},
						&models.ArticleContent{},
						&models.Tag{},
						&models.ArticleTags{},
						&models.Attachment{},
						&models.Gist{},
						&models.GistOutput{},
						&models.Comment{},
						&models.LoginRecord{},
						&models.User{},
						&models.Settings{},
					)
					if err != nil {
						return err
					}

					_, err = app.O.SqliteFtsDB.Exec(`CREATE VIRTUAL TABLE gists_fts USING fts5(
						title,
						lang,
						content,
					, content = gists,
					, content_rowid=id,
					, tokenize = 'simple')`)
					/*
						if err != nil {
							return err
						}
						// sqlite only
						_, err = app.O.SqliteFtsDB.Exec(`CREATE TRIGGER after_insert_gist
							AFTER INSERT ON gists
							BEGIN
							    INSERT INTO gists_fts (rowid, title, lang, content) VALUES (NEW.id, NEW.title, NEW.lang, NEW.content);
							END;
						`)
					*/
					return err
				},
			},
			{
				Name:  "passwd",
				Usage: "set account password",
				Action: func(ctx *cli.Context) error {
					if err := app.InitSetting(); err != nil {
						return err
					}
					if err := app.InitDatabase(); err != nil {
						return err
					}

					fmt.Print("Enter Username: ")
					reader := bufio.NewReader(os.Stdin)
					username, err := reader.ReadString('\n')
					if err != nil {
						return err
					}

					fmt.Print("Enter Password: ")
					bytePass, _ := term.ReadPassword(int(syscall.Stdin))
					pass, err := bcrypt.GenerateFromPassword(bytePass, 14)
					if err != nil {
						return err
					}

					user := models.User{
						Name: strings.TrimSpace(username),
						Pass: strings.TrimSpace(string(pass)),
					}
					// 单用户
					app.O.Where("is_del = ?", 0).Unscoped().Delete(&models.User{})
					return app.O.Create(&user).Error
				},
			},
		},
	}
	return command
}
