package core

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/go-chi/jwtauth/v5"
	"github.com/urfave/cli/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type App struct {
	isDev       bool
	RootCmd     *cli.App
	O           *gorm.DB
	Setting     *setting
	TokenAuth   *jwtauth.JWTAuth
	OssClient   *oss
	CacheClient *cacheClient
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
	// TODO sqlite
	db, err := gorm.Open(mysql.Open(app.Setting.Database.Source), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}
	app.O = db
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
	cache, err := NewCacheClient(app.Setting.Cache.Addr)
	if err != nil {
		return err
	}
	app.CacheClient = cache
	return nil
}

func (app *App) InitTokenAuth() {
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
