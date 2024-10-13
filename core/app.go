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
	Setting     *setting
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

func (app *App) InitDatabase() error {
	orm, err := newORM(app.Setting.Database.Driver, app.Setting.Database.Source)
	if err != nil {
		return err
	}
	app.O = orm
	return nil
}

func (app *App) InitOSS() error {
	oss, err := newOSS(
		app.Setting.OSS.AccessKeyId,
		app.Setting.OSS.AccessKeySecret,
		app.Setting.OSS.Region,
		app.Setting.OSS.RoleArn,
	)
	if err != nil {
		return err
	}
	app.OssClient = oss
	return nil
}

func (app *App) InitCache() error {
	cache, err := NewCache(app.Setting.Cache.Addr)
	if err != nil {
		return err
	}
	app.Cache = cache
	return nil
}

func (app *App) InitTokenAuth() {
	app.TokenAuth = jwtauth.New("HS256", []byte(app.Setting.JwtSecret), nil)
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
