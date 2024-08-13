package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

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
				Name:  "test",
				Usage: "test import",
				Action: func(ctx *cli.Context) error {
					fileContent, err := os.Open("test.json")
					if err != nil {
						return err
					}
					defer fileContent.Close()
					byteResult, _ := ioutil.ReadAll(fileContent)
					var res []map[string]interface{}
					json.Unmarshal([]byte(byteResult), &res)
					for _, v := range res {
						fmt.Printf("%s\n", v["title"])
						article := model.Article{
							Title:   v["title"].(string),
							Summary: v["summary"].(string),
						}
						article.BaseModel = model.BaseModel{
							Created: uint(v["created"].(float64)),
						}
						article.Image = v["image"].(string)
						//article.Content = v["content"].(string)
						result := app.O.Create(&article)
						println(article.ID)
						if result.Error != nil {
							panic(err)
						}
					}
					return nil
				},
			},
			{
				Name:  "construct",
				Usage: "sync database struct",
				Action: func(ctx *cli.Context) error {
					return app.O.AutoMigrate(
						&model.Article{},
						&model.ArticleText{},
						&model.ArticleTags{},
						&model.Tag{},
						&model.Attachment{},
					)
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
