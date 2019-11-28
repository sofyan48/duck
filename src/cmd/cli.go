package cmd

import (
	"github.com/urfave/cli"
)

var app *cli.App

// ArgsMapping object mapping
type ArgsMapping struct {
	ConfigPath   string
	TemplatePath string
}

// Args Glabal Acces args command
var Args ArgsMapping

// Init Initialise a CLI app
func Init() *cli.App {
	app = cli.NewApp()
	app.Name = "duck"
	app.Usage = "duck scheduler task send"
	app.Author = "meong48"
	app.Email = "meongbego@gmail.com"
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "environtmen, e",
			Usage:       "Load environtmen from `FILE`",
			Destination: &Args.ConfigPath,
		},
	}
	return app
}

// AppCommands All Command line app
func AppCommands() *cli.App {
	app := Init()
	app.Commands = []cli.Command{
		createEnvi(),
		workerStart(),
	}
	return app
}
