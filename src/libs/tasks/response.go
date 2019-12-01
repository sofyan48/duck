package tasks

import (
	"fmt"
)
// TaskResponse ...
func TaskResponse(data ...string) (string, error) {
	return getResponse(data), nil
}


func getResponse(args []string) string {
	fmt.Println(args)
	return ""
}