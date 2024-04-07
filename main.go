package main

import (
	"github.com/moonprism/blog/cmd"
	"github.com/moonprism/blog/core"
)

func main() {
	app := core.NewApp()
	app.AddSubcommand(cmd.NewServeCommand(app))
	app.AddSubcommand(cmd.NewDataCommand(app))
	app.Run()
}
