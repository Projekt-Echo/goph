package main

import (
	. "github.com/urfave/cli/v2"
	"goph/log"
	"os"
)

var AppVersion = "0.0.0"

func main() {
	app := NewApp()
	app.Name = "Goph"
	app.Description = "A Player Helper based on Golang"
	app.Usage = "Player Helper"
	app.Version = AppVersion
	app.Flags = []Flag{}
	app.Commands = []*Command{
		{
			//登录
			Name:  "login",
			Usage: "Login to get cookie",
			Action: func(context *Context) error {
				return nil
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Zapper.Panic("Cli Application Init Failed")
	}
}
