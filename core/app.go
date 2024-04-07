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
	Setting *setting
}

func (app *App) AddSubcommand(cmd *cli.Command) {
	app.RootCmd.Commands = append(app.RootCmd.Commands, cmd)
}

func (app *App) Run() error {
	if err := app.initSetting(); err != nil {
		return err
	}
	if err := app.initDB(); err != nil {
		return err
	}
	return app.RootCmd.Run(os.Args)
}

func (app *App) initSetting() error {
	setting, err := NewSetting("./app.toml")
	if err != nil {
		return err
	}
	app.Setting = &setting
	return nil
}

func (app *App) initDB() error {
	// ...
	engine, err := func(db SettingDatabase) (*xorm.Engine, error) {
		return xorm.NewEngine(db.Driver, db.Source)
	}(app.Setting.Database)
	if err != nil {
		return err
	}
	app.O = engine
	return nil
}

func NewApp() *App {
	return &App{
		isDev: false,
		RootCmd: &cli.App{
			Name: "blog",
		},
	}
}
