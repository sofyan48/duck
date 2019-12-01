package cmd

import (
	"github.com/sofyan48/duck/src/libs"
	server "github.com/sofyan48/duck/src/server/router"
	"github.com/urfave/cli"
)

// restServer mapping command
func restServer() cli.Command {
	command := cli.Command{}
	command.Name = "server"
	command.Usage = "server start, Starting REST API"
	command.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "environtment, e",
			Usage:       "Set Environtment for server",
			Destination: &Args.TemplatePath,
		},
	}
	command.Action = func(c *cli.Context) error {
		if c.Args()[0] == "start" {
			libs.LoadEnvirontment(Args.EnvPath)
			server.Init()
		}
		return nil
	}

	return command
}
