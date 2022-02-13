package cmd

import (
	/// "git.kicoe.com/blog/write/email"
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
}

func runWeb(ctx *cli.Context) error {
	Init()
	app := routers.Routers()
	app.HideBanner = true

	if !setting.App.Debug {
		app.Debug = false
		app.Use(middleware.Recover())
	}

	app.Use(middlewares.Logger)

	go RunRpc(ctx)

	// go email.RunCommentReplyNotice()

	app.Logger.Fatal(app.Start(":" + setting.App.Port))

	return nil
}
