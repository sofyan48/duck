package cmd

import (
	"os"
	"os/signal"
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
			Name:        "name",
			Usage:       "Set Worker Name",
			Destination: &Args.WorkerName,
		},

		cli.StringFlag{
			Name:        "concurrent",
			Usage:       "Set Worker Concurrent",
			Destination: &Args.WorkerConcurent,
		},
	}
	command.Action = func(c *cli.Context) error {
		if c.Args()[0] == "configure" {
			return cli.NewExitError("Under Construction", 1)
		}

		if c.Args()[0] == "start" {
			concurent, _ := strconv.Atoi(Args.WorkerConcurent)
			if concurent == 0 {
				concurent = 1
				libs.LogInfo("Load Default Concurent", 1)
			}
			libs.LogInfo("Starting Worker", "")
			workerStart(Args.EnvPath, Args.WorkerName, uint(concurent))
			libs.LogInfo("Stoped Worker", "")
		}

		return nil
	}

	return command
}

// workerStart function
// @envpath: string
func workerStart(envpath, name string, concurent uint) {
	libs.LoadEnvirontment(envpath)
	pool := libs.StartProcess(name, concurent)
	pool.Start()
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, os.Kill)
	<-signalChan
	pool.Stop()
}
