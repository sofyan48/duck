package tasks

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/RichardKnop/machinery/v1/log"
	"github.com/sofyan48/duck/src/libs/scheme"
)

func GetRequest(args []string) string {
	urls := args[0]
	methods := args[1]
	headersQuery := args[2]
	paramsQuery := args[3]
	bodyQuery := args[4]
	var body string
	if methods == "GET" {
		body = requestGet(methods, urls, headersQuery, paramsQuery)
	} else if methods == "POST" {
		body = requestPost(methods, urls, headersQuery, bodyQuery)
	}

	return string(body)
}

func requestGet(methods, urls, headers, query string) string {
	objHeaders := parsingQuery(headers)
	queryHeaders := parsingQuery(query)
	req, _ := http.NewRequest(methods, urls, nil)
	req.Header.Set("Content-Type", "application/json")
	for _, header := range objHeaders {
		req.Header.Set(header.Name, header.Value)
	}
	params := req.URL.Query()
	for _, query := range queryHeaders {
		params.Add(query.Name, query.Value)
	}
	req.URL.RawQuery = params.Encode()
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.INFO.Println("The HTTP request failed with error %s\n", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.INFO.Println("The HTTP request failed with error %s\n", err)
	}
	return string(body)
}

func requestPost(methods, urls, headers, body string) string {
	req, _ := http.NewRequest(methods, urls, bytes.NewBuffer([]byte(body)))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.INFO.Println("The HTTP request failed with error %s\n", err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.INFO.Println("The HTTP request failed with error %s\n", err)
	}
	return string(data)
}

func parsingQuery(data string) []scheme.Query {
	Query := []scheme.Query{}
	err := json.Unmarshal([]byte(data), &Query)
	if err != nil {
		log.INFO.Println("The HTTP request failed with error %s\n", err)
	}
	return Query
}
