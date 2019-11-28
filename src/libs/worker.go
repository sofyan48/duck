package libs

import (
	"github.com/RichardKnop/machinery/v1"
)

// WorkerStart Run Worker
func WorkerStart(srv *machinery.Server) error {
	// The second argument is a consumer tag
	// Ideally, each worker should have a unique tag (worker1, worker2 etc)
	worker := srv.NewWorker("machinery_worker", 0)

	if err := worker.Launch(); err != nil {
		return err
	}

	return nil
}
