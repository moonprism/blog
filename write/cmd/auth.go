package cmd

import (
	"fmt"

	"git.kicoe.com/blog/write/services/auth"
	"github.com/gookit/color"
	"github.com/urfave/cli/v2"
	"golang.org/x/crypto/ssh/terminal"
)

func upsertUser(ctx *cli.Context) error {
	var name, passwd string
	color.Println("please input username:")
	fmt.Scanf("%s", &name)
	color.Println("please input password:")
	password, _ := terminal.ReadPassword(0)
	passwd = string(password)
	err := auth.Upsert(&auth.AuthBody{
		Username: name,
		Password: passwd,
	})
	if err == nil {
		color.Println("<green>update user login info success!</>")
	}
	return err
}
