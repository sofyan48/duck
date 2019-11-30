package libs

import (
	"fmt"
	"reflect"

	"github.com/RichardKnop/machinery/v1"
	"github.com/sofyan48/duck/src/libs/scheme"
	"github.com/sofyan48/duck/src/libs/tasks"
)

// ListTask list task all
func ListTask(srv *machinery.Server, registerTask scheme.RegisterTask) {
	data := scheme.RegisterTask{}
	tasksData := make(map[string]interface{})
	for _, data := range registerTask.Worker.Register {
		tasksData[data.Action] = data.Name
		fmt.Println(tasksData)
	}
	fmt.Println(reflect.TypeOf(tasks.TaskRequest))

	tasksData1 := map[string]interface{}{
		"request": tasks.TaskRequest,
	}

	srv.RegisterTasks(tasksData1)

}
