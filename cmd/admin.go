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
	"gorm.io/gorm"
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
					if err := app.InitORM(); err != nil {
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
					if err := app.O.FtsCreateTable("gist"); err != nil {
						return err
					}
					if err := app.O.FtsCreateTable("article"); err != nil {
						return err
					}
					return nil
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
				},
			},
			{
				Name:  "sync-fts",
				Usage: "sync sqlite fts",
				Action: func(ctx *cli.Context) error {
					if err := app.InitSetting(); err != nil {
						return err
					}
					if err := app.InitORM(); err != nil {
						return err
					}
					if err := app.O.FtsCreateTable("gist"); err != nil {
						return err
					}
					if err := app.O.FtsCreateTable("article"); err != nil {
						return err
					}
					var gists []*models.Gist
					if err := app.O.Find(&gists).Error; err != nil {
						return err
					}
					for _, v := range gists {
						if err := app.O.FtsInsert("gist", v.ID, models.Gist2Text(v)); err != nil {
							return err
						}
					}
					err := app.O.Model(&models.Article{}).
						FindInBatches(&models.Article{}, 20, func(tx *gorm.DB, batch int) error {
							var articles []*models.Article
							err := tx.Preload("ArticleContent", func(db *gorm.DB) *gorm.DB {
								return db.Omit("html")
							}).Find(&articles).Error
							for _, art := range articles {
								if err := app.O.FtsInsert("article", art.ID, models.Art2Text(art)); err != nil {
									return err
								}
							}
							return err
						}).Error
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
					if err := app.InitORM(); err != nil {
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
