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
	command.Action = func(c *cli.Context) error {
		libs.LoadEnvirontment("")
		srv, err := InitServer()
		libs.Check(err)
		libs.WorkerStart(srv)
		return nil
	}
	return command
}
