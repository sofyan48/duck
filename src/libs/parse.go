package libs

import (
	"encoding/json"

	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/sofyan48/duck/src/libs/scheme"
)

// parsingArgsBody parse args body and parameter to tasks.Args
func parsingArgsBody(yamlData scheme.SendTask) string {
	bodyObj := []tasks.Arg{}
	// result := scheme.SendTaskParameter{}
	for _, body := range yamlData.Duck.Body {
		bodyData := tasks.Arg{}
		bodyData.Name = body.Name
		bodyData.Type = body.Type
		bodyData.Value = body.Value
		bodyObj = append(bodyObj, bodyData)
	}

	bodyString, _ := json.Marshal(bodyObj)
	return string(bodyString)
}

// parsingArgsParams parse args body and parameter to tasks.Args
func parsingArgsParams(yamlData scheme.SendTask) string {
	paramsObj := []tasks.Arg{}

	for _, params := range yamlData.Duck.Params {
		paramsData := tasks.Arg{}
		paramsData.Name = params.Name
		paramsData.Type = params.Type
		paramsData.Value = params.Value
		paramsObj = append(paramsObj, paramsData)
	}
	paramsString, _ := json.Marshal(paramsObj)
	return string(paramsString)
}

// parsingArgsHeaders parse args body and parameter to tasks.Args
func parsingArgsHeaders(yamlData scheme.SendTask) string {
	headersObj := []tasks.Arg{}

	for _, params := range yamlData.Duck.Headers {
		paramsData := tasks.Arg{}
		paramsData.Name = params.Name
		paramsData.Type = params.Type
		paramsData.Value = params.Value
		headersObj = append(headersObj, paramsData)
	}
	headersString, _ := json.Marshal(headersObj)
	return string(headersString)
}
