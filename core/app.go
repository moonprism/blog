package core

import (
	"os"

	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/urfave/cli/v2"
)

type App struct {
	isDev   bool
	RootCmd *cli.App
	O       *xorm.Engine
}

func (app *App) AddSubcommand(cmd *cli.Command) {
	app.RootCmd.Commands = append(app.RootCmd.Commands, cmd)
}

func (app *App) Run() error {
	app.initDB()
	return app.RootCmd.Run(os.Args)
}

func (app *App) initDB() error {
	engine, err := xorm.NewEngine("sqlite3", "./test.db")
	app.O = engine
	return err
}

func NewApp() *App {
	return &App{
		isDev: false,
		RootCmd: &cli.App{
			Name: "blog",
		},
	}
}
