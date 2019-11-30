package libs

import (
	"github.com/RichardKnop/machinery/v1"
)

// WorkerStart Run Worker
// @srv: *machinery.Server
// @name: string
// @concurency: uint
// return error
func WorkerStart(srv *machinery.Server, name string, concurency uint) error {
	worker := srv.NewWorker(name, int(concurency))
	err := worker.Launch()
	if err != nil {
		return err
	}
	return nil
}
