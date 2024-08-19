package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/moonprism/blog/core"
	"github.com/moonprism/blog/models"
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
					fileContent, err := os.Open("file.test.json")
					if err != nil {
						return err
					}
					defer fileContent.Close()
					byteResult, _ := ioutil.ReadAll(fileContent)
					var res []map[string]interface{}
					json.Unmarshal([]byte(byteResult), &res)
					// file
					for _, v := range res {
						a := models.Attachment{
							Key: v["key"].(string),
							BaseModel: models.BaseModel{
								Created: uint(v["created"].(float64)),
							},
						}
						result := app.O.Create(&a)
						println(a.ID)
						if result.Error != nil {
							panic(err)
						}
					}
					aFileContent, err := os.Open("article.test.json")
					if err != nil {
						return err
					}
					defer aFileContent.Close()
					byteResult, _ = ioutil.ReadAll(aFileContent)
					json.Unmarshal([]byte(byteResult), &res)
					// article
					for i := len(res) - 1; i >= 0; i-- {
						v := res[i]
						fmt.Printf("%s\n", v["title"])
						article := models.Article{
							Title:   v["title"].(string),
							Summary: v["summary"].(string),
						}
						article.BaseModel = models.BaseModel{
							ID:      uint(v["id"].(float64)),
							Created: uint(v["created"].(float64)),
						}
						article.Image = v["image"].(string)
						articleContent := models.ArticleContent{
							ArticleID: article.ID,
							Text:      v["content"].(string),
						}
						result := app.O.Create(&article)
						println(article.ID)
						if result.Error != nil {
							panic(err)
						}
						result = app.O.Create(&articleContent)
						if result.Error != nil {
							panic(err)
						}
					}
					// gist
					gFileContent, err := os.Open("gist.test.json")
					if err != nil {
						return err
					}
					defer gFileContent.Close()
					byteResult, _ = ioutil.ReadAll(gFileContent)
					json.Unmarshal([]byte(byteResult), &res)
					for _, v := range res {
						a := models.Gist{
							Title:   v["description"].(string),
							Lang:    v["lang"].(string),
							Content: v["content"].(string),
							BaseModel: models.BaseModel{
								Created: uint(v["created"].(float64)),
								Updated: uint(v["updated"].(float64)),
							},
						}
						result := app.O.Create(&a)
						println(a.ID)
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
						&models.Article{},
						&models.ArticleContent{},
						&models.Tag{},
						&models.ArticleTags{},
						&models.Attachment{},
						&models.Gist{},
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
