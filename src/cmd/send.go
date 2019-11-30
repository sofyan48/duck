package cmd

import (
	"fmt"
	"time"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/log"
	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/sofyan48/duck/src/libs"
	"github.com/urfave/cli"
)

func send() cli.Command {
	command := cli.Command{}
	command.Name = "send"
	command.Usage = "worker start, worker configure"
	command.Action = func(c *cli.Context) error {
		libs.LoadEnvirontment(Args.EnvPath)
		srv, err := InitServer()
		libs.Check(err)
		sendTask(srv)
		return nil
	}

	return command
}

// sendTask testt
func sendTask(server *machinery.Server) error {

	var task0 tasks.Signature
	var initTasks = func() {
		task0 = tasks.Signature{
			Name: "add",
			Args: []tasks.Arg{
				{
					Type:  "int64",
					Value: 1,
				},
				{
					Type:  "int64",
					Value: 1,
				},
			},
		}
	}

	// /*
	//  * First, let's try sending a single task
	//  */
	initTasks()
	log.INFO.Println("Single task:")

	asyncResult, err := server.SendTask(&task0)
	if err != nil {
		return fmt.Errorf("Could not send task: %s", err.Error())
	}

	results, err := asyncResult.Get(time.Duration(time.Millisecond * 5))
	if err != nil {
		return fmt.Errorf("Getting task result failed with error: %s", err.Error())
	}
	log.INFO.Printf("1 + 1 = %v\n", results[0].Interface())
	return nil
}
