package main

import (
	"encoding/json"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var client = &http.Client{Timeout: 10 * time.Second}

var gauge = prometheus.NewGauge(prometheus.GaugeOpts{
	Namespace: "prometheus_elastic_search",
	Name:      "task_current_running_time_milliseconds",
})

var pollFrequency, _ = strconv.Atoi(getEnv("ES_POLL_FREQ", "8"))
var esUrl = getEnv("ES_URL", "localhost:9200")

func getESData(url string, target interface{}) error {
	fmt.Println(url)
	r, err := client.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func collectTaskInfo() {
	taskResp := TaskResp{}
	url := "http://" + esUrl + "/_tasks?actions=*search&detailed*"

	err := getESData(url, &taskResp)
	if err != nil {
		fmt.Println("request failed: %w", err)
	}

	for key, _ := range taskResp.Nodes {
		for taskKey, _ := range taskResp.Nodes[key].Tasks {
			task := taskResp.Nodes[key].Tasks[taskKey]
			gauge.Set(float64(task.RunningTimeInNanos / 1000000))
			fmt.Printf(" \ntimestamp: %s, task_id: %d, running_time_ms: %d, "+
				"action: %s, description: %s", time.Now().Format("2006-01-02 15:04:05"), task.Id,
				task.RunningTimeInNanos/1000000, task.Action, task.Description)
		}
	}
}

func collectEvery() {
	tickEvery := time.Duration(pollFrequency) * time.Second

	for _ = range time.Tick(tickEvery) {
		collectTaskInfo()
	}
}

func main() {
	go collectEvery()
	fmt.Println("Listening")
	http.Handle("/metrics", promhttp.Handler())
	prometheus.MustRegister(gauge)
	prometheus.Unregister(prometheus.NewGoCollector())
	http.ListenAndServe(":2112", nil)
}
