package tasks

import (
	"fmt"

	"github.com/gocraft/work"
)

// GetResponse function get response from any
func GetResponse(job *work.Job) string {
	fmt.Println(job)
	return ""
}
