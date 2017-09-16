package logstash

import (
	"context"
	"encoding/json"
)

type NodeProcessStatsInfo struct {
	Process struct {
		OpenFileDescriptors     uint32 `json:"open_file_descriptors"`
		PeakOpenFileDescriptors uint32 `json:"peak_open_file_descriptors"`
		MaxFileDescriptors      uint32 `json:"max_file_descriptors"`
		Mem                     struct {
			TotalVirtualInBytes uint64 `json:"total_virtual_in_bytes"`
		} `json:"mem"`
		CPU struct {
			TotalInMillis uint32 `json:"total_in_millis"`
			Percent       uint32 `json:"percent"`
			LoadAverage   struct {
				Minitue        float64 `json:"1m"`
				FiveMinitue    float64 `json:"5m"`
				FifteenMinitue float64 `json:"15m"`
			} `json:"load_average"`
		} `json:"cpu"`
	} `json:"process"`
}

type NodeProcessStatsService struct {
	client *Client
}

const (
	node_process_stats_endpoint = "/_node/stats/process"
)

func NewNodeProcessStatsService(client *Client) *NodeProcessStatsService {
	return &NodeProcessStatsService{client: client}
}

func (n *NodeProcessStatsInfo) Json() (string, error) {
	bytes, err := json.Marshal(n)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func (n *NodeProcessStatsService) Path() string {
	return node_process_stats_endpoint
}

func (n *NodeProcessStatsService) Do(ctx context.Context) (*NodeProcessStatsInfo, error) {
	req, err := n.client.newRequest(ctx, "GET", node_process_stats_endpoint, nil)
	if err != nil {
		return nil, err
	}

	res, err := n.client.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	var np NodeProcessStatsInfo
	if err := n.client.decodeBody(res, &np); err != nil {
		return nil, err
	}

	return &np, nil
}
