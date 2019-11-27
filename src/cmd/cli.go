package cmd

import (
	"github.com/sofyan48/duck/src/configs"
	"github.com/sofyan48/duck/src/libs"
	"github.com/urfave/cli"
)

var app *cli.App

// ArgsMapping object mapping
type ArgsMapping struct {
	ConfigPath string
	TestArg    string
}

// Args Glabal Acces args command
var Args ArgsMapping

// CLIinterface obj
type CLIinterface interface {
	Init() *cli.App
}

// Init Initialise a CLI app
func Init() *cli.App {
	app = cli.NewApp()
	app.Name = "duck"
	app.Usage = "duck scheduler task send"
	app.Author = "meong"
	app.Email = "meongbego@gmail.com"
	app.Version = "0.0.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "c",
			Value:       "",
			Destination: &Args.ConfigPath,
			Usage:       "Path to a configuration file",
		},
	}
	return app
}

// AppCommands mapping command
func AppCommands() *cli.App {
	app := Init()
	app.Commands = []cli.Command{
		{
			Name:  "test",
			Usage: "launch machinery worker",
			Action: func(c *cli.Context) error {
				cfg, err := configs.LoadMachineryConfig(Args.ConfigPath)
				if err != nil {
					return cli.NewExitError(err.Error(), 1)
				}
				srv, err := libs.StartServer(cfg)
				if err != nil {
					return cli.NewExitError(err.Error(), 1)
				}
				tasks := map[string]interface{}{}
				srv.RegisterTasks(tasks)
				return nil
			},
		},
	}
	return app
}
