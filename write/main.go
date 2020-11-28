package main

import (
	"flag"
	"log"
	"os"

	"git.kicoe.com/blog/write/config"
	"git.kicoe.com/blog/write/database"
	_ "git.kicoe.com/blog/write/docs"
	"git.kicoe.com/blog/write/middlewares"
	"git.kicoe.com/blog/write/router"
	"git.kicoe.com/blog/write/workers"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Kicoe Blog API
// @version 1.0
// @description This is blog api test server. jwt token prefix: Bearer

// @BasePath /api/v1
var (
	env string
)

func init() {
	flag.StringVar(&env, "env", "dev", "environment for server:[dev|test|prod]")
	flag.Parse()
}

func main() {

	// init
	config.InitConfig(env)
	database.InitMysqlEngine()

	// todo check serve health

	go workers.RunSendEmail()

	app := router.Routers()

	if env == "prod" {
		app.Debug = false
		app.Use(middleware.Recover())
		os.MkdirAll("log", os.ModePerm)
		f, err := os.OpenFile("log/app.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer f.Close()
		app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Output: f,
		}))
	} else {
		app.Use(middlewares.Test)
		app.Use(middleware.Logger())
		app.GET("/swagger/*", echoSwagger.WrapHandler)
	}

	app.Logger.Fatal(app.Start(":" + config.App.Port))
}
