package libs

import (
	"fmt"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
)

// StartServer start server scheduler
func StartServer(cnf *config.Config) (*machinery.Server, error) {
	fmt.Println("STARTING")
	// Create server instance
	server, err := machinery.NewServer(cnf)
	if err != nil {
		return nil, err
	}

	// // Register tasks
	// tasks := map[string]interface{}{
	// 	// "add":               exampletasks.Add,
	// 	// "multiply":          exampletasks.Multiply,
	// 	// "sum_ints":          exampletasks.SumInts,
	// 	// "sum_floats":        exampletasks.SumFloats,
	// 	// "concat":            exampletasks.Concat,
	// 	// "split":             exampletasks.Split,
	// 	// "panic_task":        exampletasks.PanicTask,
	// 	// "long_running_task": exampletasks.LongRunningTask,
	// }

	return server, nil
}
