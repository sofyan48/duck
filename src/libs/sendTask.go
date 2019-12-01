package libs

import (
	"errors"
	"fmt"
	"time"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/log"
	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/sofyan48/duck/src/libs/scheme"
)

// SendTask testt
func SendTask(server *machinery.Server, yamlData scheme.SendTask) (interface{}, error) {
	var task0 tasks.Signature
	if yamlData.Duck.Action.Trigger != "request" {
		fmt.Println(yamlData.Duck.Action.Trigger)
		err := errors.New("Trigerr response Or request")
		return nil, err
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
					Value: "",
					// Value: libs.parsingArgs(yamlData),
					Name: "body",
				},
			},
		}
		// eta := time.Now().UTC().Add(time.Second * 20)
		// task0.ETA = &eta
		task0.IgnoreWhenTaskNotRegistered = true
	}

	initTasks()
	log.INFO.Println("Send task:", yamlData.Duck.Task)
	asyncResult, err := server.SendTask(&task0)
	if err != nil {
		return nil, fmt.Errorf("Could not send task: %s", err.Error())
	}
	taskState := asyncResult.GetState()
	results, err := asyncResult.Get(time.Duration(time.Millisecond * 5))
	if err != nil {
		return nil, fmt.Errorf("Getting task result failed with error: %s", err.Error())
	}
	log.INFO.Printf("UUID = %s", taskState.TaskUUID)
	log.INFO.Printf("Results = %v", results[0].Interface())
	response := map[string]interface{}{
		"Result": results[0].Interface(),
		"UUID":   taskState.TaskUUID,
	}
	return response, nil
}
