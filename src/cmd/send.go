package cmd

import (
	"fmt"

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
			templates = libs.GetPCurrentPath() + "/send.yml"
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
		fmt.Println(ymlData)
		return nil
	}

	return command
}
