package main

type DetailedNode struct {
	Node               string `json:"node"`
	Id                 int    `json:"id"`
	Type               string `json:"type"`
	Action             string `json:"action"`
	Description        string `json:"description"`
	StartTimeInMillis  int64  `json:"start_time_in_millis"`
	RunningTimeInNanos int    `json:"running_time_in_nanos"`
	Cancellable        bool   `json:"cancellable"`
	Cancelled          bool   `json:"cancelled"`
}

type BasicNode struct {
	Name             string `json:"name"`
	TransportAddress string `json:"transport_address"`
	Host             string `json:"host"`
	Ip               string `json:"ip"`
	Tasks            map[string]*DetailedNode
}

type TaskResp struct {
	Nodes map[string]*BasicNode
}
