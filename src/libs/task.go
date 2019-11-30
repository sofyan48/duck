package libs

import (
	"github.com/RichardKnop/machinery/v1"
	"github.com/sofyan48/duck/src/libs/tasks"
)

// ListTask list task all
func ListTask(srv *machinery.Server) {
	// list All Data
	tasksData := map[string]interface{}{
		"add": tasks.Add,
	}

	srv.RegisterTasks(tasksData)

}
