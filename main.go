package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var client = &http.Client{Timeout: 10 * time.Second}

func getESData(url string, target interface{}) error {
	r, err := client.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

func collectTaskInfo() {
	taskResp := TaskResp{}

	err := getESData("http://localhost:3000/_tasks?actions=*search&detailed*", &taskResp)
	if err != nil {
		fmt.Println("request failed: %w", err)
	}

	for key, _ := range taskResp.Nodes {
		for taskKey, _ := range taskResp.Nodes[key].Tasks {
			task := taskResp.Nodes[key].Tasks[taskKey]
			prom

		}
	}
}

func main() {
	collectTaskInfo()
	fmt.Println("Listening")
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
