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

	ind := se.NewIndex("code")
	for i := 0; ; i++ {
		codes, _, _ := code.GetList(i, 100)
		codeSlice := []map[string]interface{}{}
		for _, c := range codes {
			color.Println("<green>" + c.Description + "</>")
			codeDetail, _ := code.GetDetail(c.ID)
			color.Println("<yellow>" + codeDetail.Content + "</>")
			codeSlice = append(codeSlice, code.ToDoc(codeDetail.Code))
		}
		ind.Insert(codeSlice)
		if len(codes) < 100 {
			break
		}
	}
	return nil
}
