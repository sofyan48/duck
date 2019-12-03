package libs

import (
	"fmt"

	"github.com/gocraft/work"
	"github.com/sofyan48/duck/src/config"
)

// StartProcess Start Process Server
func StartProcess(processName string, concurency uint) *work.WorkerPool {
	redisPool := config.LoadConfig()
	pool := work.NewWorkerPool(ContextTask{}, concurency, processName, redisPool)
	pool.Middleware((*ContextTask).Log)
	ListTask(pool)
	return pool
}

// Log middleware for log job task
func (ctx *ContextTask) Log(job *work.Job, next work.NextMiddlewareFunc) error {
	LogInfo(job.Name)
	fmt.Println("Starting job: ", job.Name)
	return next()
}
