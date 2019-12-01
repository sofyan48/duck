package cmd

import (
	"github.com/urfave/cli"
)

var app *cli.App

// ArgsMapping object mapping
type ArgsMapping struct {
	EnvPath         string
	TemplatePath    string
	WorkerName      string
	WorkerConcurent string
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
			Name:        "config, c",
			Usage:       "Load environtment config from `FILE`",
			Destination: &Args.EnvPath,
		},
	}
	return app
}

// AppCommands All Command line app
func AppCommands() *cli.App {
	app := Init()
	app.Commands = []cli.Command{
		createEnvi(),
		worker(),
		send(),
		restServer(),
	}
	return app
}
