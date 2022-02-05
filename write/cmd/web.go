package cmd

import (
	"git.kicoe.com/blog/write/email"
	"git.kicoe.com/blog/write/modules/se"
	"git.kicoe.com/blog/write/modules/setting"
	"git.kicoe.com/blog/write/web/middlewares"
	"git.kicoe.com/blog/write/web/routers"
	"github.com/labstack/echo/v4/middleware"
	"github.com/urfave/cli/v2"
)

var CmdWeb = cli.Command{
	Name:        "web",
	Usage:       "start web server",
	Description: "web server need run",
	Action:      runWeb,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "env",
			Value: "dev",
			Usage: "dev/prod configuration file name",
		},
		&cli.BoolFlag{
			Name: "reindex",
			Value: false,
			Usage: "reindex search data",
		},
		&cli.BoolFlag{
			Name: "passwd",
			Value: false,
			Usage: "insert or update user login info",
		},
	},
}

func runWeb(ctx *cli.Context) error {
	env := ctx.String("env")
	Init(env)

	if ctx.Bool("passwd") {
		upsertUser(ctx)
		return nil
	}

	app := routers.Routers()
	app.HideBanner = true
	if env == "prod" {
		app.Debug = false
		app.Use(middleware.Recover())
	} else {

	}

	// 文件日志
	app.Use(middlewares.Logger)

	se.Init()
	go RunRpc(ctx)

	go email.RunCommentReplyNotice()

	app.Logger.Fatal(app.Start(":" + setting.App.Port))

	return nil
}
