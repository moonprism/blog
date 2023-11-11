package cmd

import (
	"strconv"

	"git.kicoe.com/blog/write/models"
	"git.kicoe.com/blog/write/modules/db"
	"github.com/gookit/color"
	"github.com/urfave/cli/v2"
)

var CmdRune = cli.Command{
	Name:        "rune",
	Usage:       "sync rune field",
	Description: "",
	Subcommands: []*cli.Command{
		{
			Name:   "sync",
			Usage:  "sync rune field",
			Action: syncRuneField,
		},
	},
}

func syncRuneField(ctx *cli.Context) error {
	Init()
	articles := make([]*models.Article, 0)
	err := db.MysqlXEngine.Table("article").Limit(1000).Find(&articles)
	if err != nil {
		return err
	}
	for _, a := range articles {
		models.UpdateArticleRune(a)
		_, err := db.MysqlXEngine.Id(a.ID).Update(a)
		if err != nil {
			return err
		}
		color.Println("<red>" + a.Title + strconv.FormatInt(a.Rune, 10) + "</>")
	}
	return nil
}
