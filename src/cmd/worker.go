package cmd

import (
	"os"
	"strconv"

	"github.com/sofyan48/duck/src/libs"
	"github.com/urfave/cli"
)

// workerStart mapping command
func worker() cli.Command {
	command := cli.Command{}
	command.Name = "worker"
	command.Usage = "worker start, worker configure"
	command.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "file, f",
			Usage:       "Load YAML configuration from `FILE`",
			Destination: &Args.TemplatePath,
		},

		cli.StringFlag{
			Name:        "name",
			Usage:       "Set Worker Name",
			Destination: &Args.WorkerName,
		},

		cli.StringFlag{
			Name:        "concurent",
			Usage:       "Set Worker Concurent",
			Destination: &Args.WorkerConcurent,
		},
	}
	command.Action = func(c *cli.Context) error {
		if c.Args()[0] == "configure" {
			os.Exit(0)
		}

		if c.Args()[0] == "start" {
			concurent, _ := strconv.Atoi(Args.WorkerConcurent)

			workerStart(Args.EnvPath, Args.WorkerName, uint(concurent))
		}

		return nil
	}

	return command
}

// workerStart function
// @envpath: string
func workerStart(envpath, name string, concurent uint) {
	libs.LoadEnvirontment(envpath)
	srv, err := InitServer()
	libs.Check(err)
	libs.ListTask(srv)
	err = libs.WorkerStart(srv, name, concurent)
	libs.Check(err)

}
