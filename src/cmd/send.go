package cmd

import (
	"github.com/sofyan48/duck/src/libs"
	"github.com/urfave/cli"
)

func send() cli.Command {
	command := cli.Command{}
	command.Name = "send"
	command.Usage = "worker start, worker configure"
	command.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "file, f",
			Usage:       "Set Action Sned From Yaml File",
			Destination: &Args.TemplatePath,
		},
	}
	command.Action = func(c *cli.Context) error {
		// init template file
		argsFile := Args.TemplatePath
		var templates string
		if argsFile == "" {
			templates = libs.GetPCurrentPath() + "/duck.yml"
		} else {
			templates = argsFile
		}
		if !libs.CheckFile(templates) {
			return cli.NewExitError("No Templates For Send", 1)
		}
		ymlData, err := libs.ReadYMLSend(templates)
		libs.Check(err)

		//load environtment
		libs.LoadEnvirontment(Args.EnvPath)
		result, err := libs.Send(ymlData.Duck.Action.Worker, ymlData.Duck.Action.Trigger, ymlData)
		if err != nil {
			libs.LogFatal("Fatal: ", err)
		}
		libs.LogInfo("Starting Queue : ", result.QueueName)
		libs.LogInfo("UUID : ", result.UUID)
		libs.LogInfo("Task : ", result.TaskName)
		libs.LogInfo("Created : ", result.CreatedAt)
		return nil
	}

	return command
}
