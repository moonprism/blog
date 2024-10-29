package core

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/go-chi/jwtauth/v5"
	"github.com/urfave/cli/v2"
)

type App struct {
	RootCmd     *cli.App
	O           *orm
	Settings    *settings
	TokenAuth   *jwtauth.JWTAuth
	OssClient   *oss
	Cache       Cache
	TmplManager *tmplManager
	Validator   *Validator
}

func (app *App) AddSubcommand(cmd *cli.Command) {
	app.RootCmd.Commands = append(app.RootCmd.Commands, cmd)
}

func (app *App) Run() error {
	return app.RootCmd.Run(os.Args)
}

func (app *App) InitSetting() error {
	settings, err := NewSettings()
	if err != nil {
		return err
	}
	app.Settings = &settings
	return nil
}

func (app *App) InitORM() error {
	orm, err := newORM(
		app.Settings.Database.Driver,
		app.Settings.Database.Source,
		app.Settings.FTS.Source,
	)
	if err != nil {
		return err
	}
	app.O = orm
	return nil
}

func (app *App) InitOSS() error {
	oss, err := newOSS(
		app.Settings.OSS.AccessKeyId,
		app.Settings.OSS.AccessKeySecret,
		app.Settings.OSS.Region,
		app.Settings.OSS.RoleArn,
	)
	if err != nil {
		return err
	}
	app.OssClient = oss
	return nil
}

func (app *App) InitCache() error {
	cache, err := NewCache(app.Settings.Cache.Addr)
	if err != nil {
		return err
	}
	app.Cache = cache
	return nil
}

func (app *App) InitTokenAuth() {
	app.TokenAuth = jwtauth.New("HS256", []byte(app.Settings.JwtSecret), nil)
}

func (app *App) InitTmpl() error {
	app.TmplManager = NewTmplManager()
	return app.TmplManager.RegisterDir("./ui/vanilla/dist")
}

func (app *App) InitValidator() {
	app.Validator = NewValidator()
}

func (app *App) JSON(w http.ResponseWriter, data any) error {
	return json.NewEncoder(w).Encode(data)
}

func (app *App) HTML(w http.ResponseWriter, tmplName string, data any) error {
	return app.TmplManager.Execute(tmplName, w, data)
}

func NewApp() *App {
	return &App{
		RootCmd: &cli.App{
			Name: "blog",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "conf",
					Aliases:     []string{"c"},
					Value:       configPath,
					Usage:       "app config path",
					Destination: &configPath,
				},
			},
		},
	}
}
