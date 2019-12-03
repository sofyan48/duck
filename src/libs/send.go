package libs

import (
	"encoding/json"

	"github.com/gocraft/work"
	"github.com/sofyan48/duck/src/config"
	"github.com/sofyan48/duck/src/libs/scheme"
)

// Send function to queue engine
// @enqName: string
// @jobName: string
func Send(enqName string, jobName string, yamlData scheme.SendTask) (scheme.SendResponse, error) {
	result := scheme.SendResponse{}
	redisPool := config.LoadConfig()
	var enqueuer = work.NewEnqueuer(enqName, redisPool)
	action, _ := json.Marshal(yamlData.Duck.Action)
	headers, _ := json.Marshal(yamlData.Duck.Headers)
	body, _ := json.Marshal(yamlData.Duck.Body)
	params, _ := json.Marshal(yamlData.Duck.Params)
	data, err := enqueuer.Enqueue(jobName, work.Q{
		"action":    string(action),
		"headers":   string(headers),
		"parameter": string(params),
		"body":      string(body),
	})
	if err != nil {
		Check(err)
		return result, err
	}
	result.UUID = data.ID
	result.TaskName = data.Name
	result.QueueName = enqName
	result.CreatedAt = ConvertUnixTime(data.EnqueuedAt)
	return result, nil
}
