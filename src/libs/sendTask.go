package libs

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/log"
	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/sofyan48/duck/src/libs/scheme"
)

// SendTask testt
func SendTask(server *machinery.Server, yamlData scheme.SendTask) (scheme.ResultsSend, error) {
	response := scheme.ResultsSend{}
	var task0 tasks.Signature
	if yamlData.Duck.Action.Trigger != "request" {
		fmt.Println(yamlData.Duck.Action.Trigger)
		err := errors.New("Trigerr response Or request")
		return response, err
	}

	var initTasks = func() {
		bodyQuery := parsingArgsBody(yamlData)
		paramsQuery := parsingArgsParams(yamlData)
		headersQuery := parsingArgsHeaders(yamlData)
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
					Value: yamlData.Duck.Action.Method,
					Name:  "method",
				},
				{
					Type:  "string",
					Value: headersQuery,
					Name:  "headers",
				},
				{
					Type:  "string",
					Value: paramsQuery,
					Name:  "parameters",
				},
				{
					Type:  "string",
					Value: bodyQuery,
					Name:  "body",
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
		return response, errors.New(err.Error())
	}
	taskState := asyncResult.GetState()
	results, err := asyncResult.Get(time.Duration(time.Millisecond * 5))
	if err != nil {
		return response, errors.New(err.Error())
	}
	log.INFO.Printf("UUID = %s", taskState.TaskUUID)
	log.INFO.Printf("Results = %v", results[0].String())
	data := make(map[string]interface{})
	err = json.Unmarshal([]byte(results[0].String()), &data)
	if err != nil {
		return response, errors.New(err.Error())
	}
	response.Result = data
	response.UUID = taskState.TaskUUID
	return response, nil
}
