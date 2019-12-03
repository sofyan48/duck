package libs

import (
	"fmt"

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
	results, err := tasks.GetRequest(job)
	if err != nil {
		return err
	}
	fmt.Println(results)
	return nil
}

// TaskResponse ...
func (ctx *ContextTask) TaskResponse(job *work.Job) error {
	// return tasks.GetResponse(string(job.Args)), nil
	return nil
}
