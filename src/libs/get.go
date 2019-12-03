package libs

import (
	"encoding/json"

	"github.com/sofyan48/duck/src/config"
)

// GetResult get reslut from jobs
func GetResult(worker, jobsName, uuid string) (map[string]interface{}, error) {
	config.GetConnection()
	keys := worker + ":" + jobsName + ":" + uuid
	resultRequest, err := config.GetRowsCached(keys)
	if err != nil {
		LogFatal("Error : ", err)
		return nil, err
	}
	result := make(map[string]interface{})
	err = json.Unmarshal([]byte(resultRequest), &result)
	if err != nil {
		LogFatal("Error : ", err)
		return nil, err
	}
	return result, nil
}
