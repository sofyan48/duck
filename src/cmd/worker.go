package cmd

import (
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
			Usage:       "Set Register Task From Yaml File",
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
			return cli.NewExitError("Under Construction", 1)
		}

		if c.Args()[0] == "start" {
			argsFile := Args.TemplatePath
			var templates string
			if argsFile == "" {
				templates = libs.GetPCurrentPath() + "/worker.yml"
			} else {
				templates = argsFile
			}
			if !libs.CheckFile(templates) {
				return cli.NewExitError("No Worker Register", 1)
			}
			concurent, _ := strconv.Atoi(Args.WorkerConcurent)
			workerStart(Args.EnvPath, Args.WorkerName, uint(concurent), templates)
		}

		return nil
	}

	return command
}

// workerStart function
// @envpath: string
func workerStart(envpath, name string, concurent uint, template string) {
	libs.LoadEnvirontment(envpath)
	srv, err := libs.InitServer(envpath)
	libs.Check(err)
	ymlData, err := libs.ReadYML(template)
	libs.ListTask(srv, ymlData)
	if name == "" {
		name = "duck"
	}
	err = libs.WorkerStart(srv, name, concurent)
	libs.Check(err)

}
