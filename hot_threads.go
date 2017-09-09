package logstash

import (
	"context"
	"encoding/json"
)

type HotThreadsInfo struct {
	HotThreads struct {
		Time           string   `json:"time"`
		BusiestThreads int      `json:"busiest_threads"`
		Threads        []Thread `json:"threads"`
	} `json:"hot_threads"`
}

type Thread struct {
	Name             string   `json:"name"`
	ThreadID         int      `json:"thread_id"`
	PercentOfCPUTime float64  `json:"percent_of_cpu_time"`
	State            string   `json:"state"`
	Traces           []string `json:"traces"`
}

type HotThreadsInfoService struct {
	client *Client
}

const (
	threads_info_endpoint = "/_node/hot_threads"
)

func NewHotThreadsInfoService(client *Client) *HotThreadsInfoService {
	return &HotThreadsInfoService{client: client}
}

func (t *HotThreadsInfoService) Path() string {
	return threads_info_endpoint
}

func (t *HotThreadsInfo) Json() (string, error) {
	bytes, err := json.Marshal(t)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func (t *HotThreadsInfoService) Do(ctx context.Context) (*HotThreadsInfo, error) {
	req, err := t.client.newRequest(ctx, "GET", threads_info_endpoint, nil)
	if err != nil {
		return nil, err
	}

	res, err := t.client.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	var ht HotThreadsInfo
	if err := t.client.decodeBody(res, &ht); err != nil {
		return nil, err
	}

	return &ht, nil
}
