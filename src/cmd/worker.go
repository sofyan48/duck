package cmd

import (
	"github.com/sofyan48/duck/src/libs"
	"github.com/urfave/cli"
)

// workerStart mapping command
func workerStart() cli.Command {
	command := cli.Command{}
	command.Name = "worker"
	command.Usage = "worker"
	command.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "yml, y",
			Usage:       "Load YAML configuration from `FILE`",
			Destination: &Args.TemplatePath,
		},
	}
	command.Action = func(c *cli.Context) error {
		libs.LoadEnvirontment(Args.ConfigPath)
		srv, err := InitServer()
		libs.Check(err)
		libs.WorkerStart(srv)
		return nil
	}
	return command
}
