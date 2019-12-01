package libs

import (
	"encoding/json"

	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/sofyan48/duck/src/libs/scheme"
)

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
