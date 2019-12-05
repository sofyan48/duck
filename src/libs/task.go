package libs

import (
	"github.com/gocraft/work"
	"github.com/sofyan48/duck/src/libs/tasks"
)

// ContextTask ...
type ContextTask struct {
}

// ListTask list task all
func ListTask(pool *work.WorkerPool) {
	pool.Job("request", (*ContextTask).TaskRequest)
	pool.Job("response", (*ContextTask).TaskResponse)
}

// TaskRequest ...
func (ctx *ContextTask) TaskRequest(job *work.Job) error {
	data, err := tasks.GetRequest(job)
	Check(err)
	results, err := ParseJSON(data)
	if err != nil {
		return err
	}
	for i, elem := range results {
		LogInfo(i+" : ", elem)
	}
	return nil
}

// TaskResponse ...
func (ctx *ContextTask) TaskResponse(job *work.Job) error {
	// return tasks.GetResponse(string(job.Args)), nil
	return nil
}
