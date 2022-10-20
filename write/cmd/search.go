package cmd

import (
	"git.kicoe.com/blog/write/modules/se"
	"git.kicoe.com/blog/write/services/code"
	"github.com/gookit/color"
	"github.com/urfave/cli/v2"
)

var CmdSE = cli.Command{
	Name:        "se",
	Usage:       "== Search Engine ==, meilisearch",
	Description: "",
	Subcommands: []*cli.Command{
		{
			Name:   "reindex",
			Usage:  "reindex from db",
			Action: seReIndex,
		},
	},
}

func seReIndex(ctx *cli.Context) error {
	Init()

	for i := 0; ; i++ {
		codes, _, _ := code.GetList(i, 100)
		for _, c := range codes {
			color.Println("<green>" + c.Description + "</>")
			codeDetail, _ := code.GetDetail(c.ID)
			se.Engine.Insert("code", codeDetail.ID, code.ToSearchDoc(codeDetail.Code))
		}
		if len(codes) < 100 {
			break
		}
	}

	return nil
}
