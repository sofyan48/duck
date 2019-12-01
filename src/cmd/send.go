package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/log"
	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/sofyan48/duck/src/libs"
	"github.com/sofyan48/duck/src/libs/scheme"
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
		// init server
		srv, err := InitServer()
		libs.Check(err)
		// send task
		err = sendTask(srv, ymlData)
		if err != nil {
			return cli.NewExitError(err.Error(), 1)
		}
		return nil
	}

	return command
}

// sendTask testt
func sendTask(server *machinery.Server, yamlData scheme.SendTask) error {
	var task0 tasks.Signature
	if yamlData.Duck.Action.Trigger != "request" {
		fmt.Println(yamlData.Duck.Action.Trigger)
		err := errors.New("Trigerr response Or request")
		return err
	}

	var initTasks = func() {
		task0 = tasks.Signature{
			Name: yamlData.Duck.Action.Trigger,
			Args: []tasks.Arg{
				{
					Type:  "string",
					Value: yamlData.Duck.Action.URL,
					Name:  "url",
				},
				{
					Type:  "string",
					Value: parsingArgs(yamlData),
					Name:  "body",
				},
			},
		}
		// eta := time.Now().UTC().Add(time.Second * 20)
		// task0.ETA = &eta
		task0.IgnoreWhenTaskNotRegistered = true
	}

	// /*
	//  * First, let's try sending a single task
	//  */
	initTasks()
	log.INFO.Println("Send task:", yamlData.Duck.Task)
	asyncResult, err := server.SendTask(&task0)
	if err != nil {
		return fmt.Errorf("Could not send task: %s", err.Error())
	}

	results, err := asyncResult.Get(time.Duration(time.Millisecond * 5))
	if err != nil {
		return fmt.Errorf("Getting task result failed with error: %s", err.Error())
	}
	log.INFO.Printf("Results = %v\n", results[0].Interface())
	return nil
}

// parsingArgs parse args body and parameter to tasks.Args
func parsingArgs(yamlData scheme.SendTask) string {
	paramsObj := []tasks.Arg{}
	for _, body := range yamlData.Duck.Body {
		bodyData := tasks.Arg{}
		bodyData.Name = body.Name
		bodyData.Type = body.Type
		bodyData.Value = body.Value
		paramsObj = append(paramsObj, bodyData)
	}
	for _, params := range yamlData.Duck.Params {
		paramsData := tasks.Arg{}
		paramsData.Name = params.Name
		paramsData.Type = params.Type
		paramsData.Value = params.Value
		paramsObj = append(paramsObj, paramsData)
	}

	results, _ := json.Marshal(paramsObj)
	return string(results)
}
