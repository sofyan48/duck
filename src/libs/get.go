package libs

import (
	"encoding/json"

	"github.com/sofyan48/duck/src/config"
	"github.com/sofyan48/duck/src/libs/redis"
)

// GetResult get reslut from jobs
func GetResult(worker, jobsName, uuid string) (map[string]interface{}, error) {
	cnf := config.Config{}
	dial, _ := cnf.LoadBroker(2)
	store := dial.RedisDial
	rds := redis.Redis{}
	keys := worker + ":" + jobsName + ":" + uuid
	resultRequest, err := rds.GetRowsCached(store, keys)
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
