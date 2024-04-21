package core

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/go-chi/jwtauth/v5"
	"github.com/urfave/cli/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type App struct {
	isDev     bool
	RootCmd   *cli.App
	O         *gorm.DB
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
	if err := app.initDatabase(); err != nil {
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

func (app *App) initDatabase() error {
	// TODO sqlite
	db, err := gorm.Open(mysql.Open(app.Setting.Database.Source), &gorm.Config{})
	if err != nil {
		return err
	}
	app.O = db
	return nil
}

func (app *App) initTokenAuth() {
	app.TokenAuth = jwtauth.New("HS256", []byte(app.Setting.JwtSecret), nil)
}

func (app *App) JSON(w http.ResponseWriter, data any) error {
	return json.NewEncoder(w).Encode(data)
}

func NewApp() *App {
	return &App{
		isDev: false,
		RootCmd: &cli.App{
			Name: "blog",
		},
	}
}
