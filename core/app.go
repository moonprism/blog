package core

import (
	"os"

	"github.com/go-chi/jwtauth/v5"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/urfave/cli/v2"
)

type App struct {
	isDev     bool
	RootCmd   *cli.App
	O         *xorm.Engine
	Setting   *setting
	TokenAuth *jwtauth.JWTAuth
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
	app.initTokenAuth()
	return app.RootCmd.Run(os.Args)
}

func (app *App) initSetting() error {
	setting, err := NewSetting()
	if err != nil {
		return err
	}
	app.Setting = &setting
	return nil
}

func (app *App) ReSetting() error {
	return ReSetting(app.Setting)
}

func (app *App) initDB() error {
	// ...
	engine, err := func(db settingDatabase) (*xorm.Engine, error) {
		return xorm.NewEngine(db.Driver, db.Source)
	}(app.Setting.Database)
	if err != nil {
		return err
	}
	engine.ShowSQL(true)
	app.O = engine
	return nil
}

func (app *App) initTokenAuth() {
	app.TokenAuth = jwtauth.New("HS256", []byte(app.Setting.JwtSecret), nil)
}

func NewApp() *App {
	return &App{
		isDev: false,
		RootCmd: &cli.App{
			Name: "blog",
		},
	}
}
