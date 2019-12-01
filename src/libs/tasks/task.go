package tasks

import "fmt"

// TaskRequest ...
func TaskRequest(url ...string) (string, error) {
	fmt.Println(url)
	return getRequest(url), nil
}

func getRequest(args []string) string {
	// response, err := http.Get(args[0])
	// if err != nil {
	// 	log.INFO.Println("The HTTP request failed with error %s\n", err)
	// }

	// data, _ := ioutil.ReadAll(response.Body)

	// return string(data)
	return ""
}
