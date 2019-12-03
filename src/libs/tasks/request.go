package tasks

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gocraft/work"
	"github.com/sofyan48/duck/src/libs/scheme"
)

// GetRequest function get request from queue
// @job: *work.Job
func GetRequest(job *work.Job) (string, error) {
	action := &scheme.ActionScheme{}
	body := job.ArgString("body")
	params := job.ArgString("parameter")
	headers := job.ArgString("headers")
	err := json.Unmarshal([]byte(job.ArgString("action")), action)
	if err != nil {
		log.Panic("ERROR :", err)
		return "", err
	}

	if action.Method == "GET" {
		body = requestGet(action.Method, action.URL, headers, params)
	} else if action.Method == "POST" {
		body = requestPost(action.Method, action.URL, headers, params)
	}

	return string(body), nil
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
		log.Println("The HTTP request failed with error", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("The HTTP request failed with error ", err)
	}
	return string(body)
}

func requestPost(methods, urls, headers, body string) string {
	req, _ := http.NewRequest(methods, urls, bytes.NewBuffer([]byte(body)))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("The HTTP request failed with error ", err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("The HTTP request failed with error ", err)
	}
	return string(data)
}

func parsingQuery(data string) []scheme.Query {
	Query := []scheme.Query{}
	err := json.Unmarshal([]byte(data), &Query)
	if err != nil {
		log.Println("The HTTP request failed with error ", err)
	}
	return Query
}
