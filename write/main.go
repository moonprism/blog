package main

import (
	"os"

	"git.kicoe.com/blog/write/cmd"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

// @title Kicoe Blog API
// @version 1.0
// @description This is blog api test server. jwt token prefix: Bearer

// @BasePath /api/v1
func main() {
	app := &cli.App{
		Name:  "write",
		Usage: "blog admin api",
		Commands: []*cli.Command{
			&cmd.CmdWeb,
			&cmd.CmdAuth,
			&cmd.CmdSE,
			&cmd.CmdSync,
			&cmd.CmdRune,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal("Failed to run app with %s: %v", os.Args, err)
	}
}
