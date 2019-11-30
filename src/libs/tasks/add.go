package tasks

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/RichardKnop/machinery/v1/log"
)

// TaskRequest ...
func TaskRequest(url ...string) (string, error) {
	return getRequest(url), nil
}

func getRequest(args []string) string {
	fmt.Println(args[0])
	response, err := http.Get(args[0])
	if err != nil {
		log.INFO.Println("The HTTP request failed with error %s\n", err)
	}

	data, _ := ioutil.ReadAll(response.Body)

	return string(data)
}
