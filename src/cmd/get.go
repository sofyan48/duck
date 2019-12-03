package cmd

import (
	"github.com/sofyan48/duck/src/libs"
	"github.com/urfave/cli"
)

func get() cli.Command {
	command := cli.Command{}
	command.Name = "get"
	command.Usage = "get Results"
	command.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "file, f",
			Usage:       "Set Action Sned From Yaml File",
			Destination: &Args.TemplatePath,
		},

		cli.StringFlag{
			Name:        "uuid, i",
			Usage:       "Set UUID RESULT",
			Destination: &Args.UUID,
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
		results, _ := libs.GetResult(ymlData.Duck.Action.Worker, ymlData.Duck.Action.Trigger, Args.UUID)
		for i, element := range results {
			libs.LogInfo(i+" : ", element)
		}
		return nil
	}

	return command
}
