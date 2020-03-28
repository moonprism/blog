package main

import (
	"flag"
	"git.kicoe.com/blog/write/config"
	"git.kicoe.com/blog/write/database"
	_ "git.kicoe.com/blog/write/docs"
	"git.kicoe.com/blog/write/router"
	"git.kicoe.com/blog/write/utils"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Kicoe Blog API
// @version 1.0
// @description This is blog api test server. jwt token prefix: Bearer

// @BasePath /api/v1
var (
	env	string
)

func init() {
	flag.StringVar(&env, "env", "dev", "environment for server:[dev|test|prod]")
	flag.Parse()
}

func main() {
	// init
	config.InitConfig(env)
	database.InitMysqlEngine()
	if env == "prod" {
		utils.InitEs()
		utils.InitOss()
	}

	app := router.Routers()
	app.GET("/swagger/*", echoSwagger.WrapHandler)
	app.Logger.Fatal(app.Start(":"+config.App.Port))
}
