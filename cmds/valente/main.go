package main

import (
	"bytes"
	"log"
	"os"

	"github.com/codegangsta/cli"
	"github.com/dimiro1/banner"
	"github.com/mattn/go-colorable"
)

const textBanner = `
__     __    _            _       
\ \   / /_ _| | ___ _ __ | |_ ___ 
 \ \ / / _' | |/ _ \ '_ \| __/ _ \
  \ V / (_| | |  __/ | | | ||  __/
   \_/ \__,_|_|\___|_| |_|\__\___|

GoVersion: {{ .GoVersion }}
GOOS: {{ .GOOS }}

`

func main() {
	isEnabled := true
	isColorEnabled := true
	banner.Init(colorable.NewColorableStdout(), isEnabled, isColorEnabled, bytes.NewBufferString(textBanner))

	trumae := cli.Author{Name: "Trumae da Ilha", Email: "trumae@gmail.com"}
	app := cli.NewApp()
	app.Name = "valente"
	app.Version = "0.0.1"
	app.Authors = []cli.Author{trumae}
	app.Usage = "Tool for easy use of valente websocket micro-framework"

	app.Commands = []cli.Command{
		{
			Name:    "new",
			Aliases: []string{"n"},
			Usage:   "create a new project",
			Action: func(c *cli.Context) {
				appname := c.Args().First()
				log.Println("Creating app ", appname)
				createApp(appname)
			},
		},
		{
			Name:    "form",
			Aliases: []string{"f"},
			Usage:   "create a simple form",
			Action: func(c *cli.Context) {
				formname := c.Args().First()
				log.Println("Creating form ", formname)
				createForm(formname)
			},
		},
		/*
			{
				Name:    "template",
				Aliases: []string{"r"},
				Usage:   "options for task templates",
				Subcommands: []cli.Command{
					{
						Name:  "add",
						Usage: "add a new template",
						Action: func(c *cli.Context) {
							println("new task template: ", c.Args().First())
						},
					},
					{
						Name:  "remove",
						Usage: "remove an existing template",
						Action: func(c *cli.Context) {
							println("removed task template: ", c.Args().First())
						},
					},
				},
			},*/
	}

	app.Run(os.Args)
}
