package logstash

import (
	"context"
	"encoding/json"
)

type NodeStatsOSInfo struct {
	OS struct {
		CGroup struct {
			CPUAcct struct {
				UsageNanos   uint64 `json:"usage_nanos"`
				ControlGroup string `json:"control_group"`
			} `json:"cpuacct"`
			CPU struct {
				CfsQuotaMicros int    `json:"cfs_quota_micros"`
				ControlGroup   string `json:"control_group"`
				Stat           struct {
					NumberOfTimesThrottled uint64 `json:"number_of_times_throttled"`
					TimeThrottledNanos     uint64 `json:"time_throttled_nanos"`
					NumberOfElapsedPeriods uint64 `json:"number_of_elapsed_periods"`
				} `json:"stat"`
				CfsPeriosMicros int `json:"cfs_period_micros"`
			} `json:"cpu"`
		} `json:"cgroup"`
	} `json:"os"`
}

type NodeStatsOSService struct {
	client *Client
}

const (
	node_stats_os_endpoint = "/_node/stats/os"
)

func NewNodeStatsOSService(client *Client) *NodeStatsOSService {
	return &NodeStatsOSService{client: client}
}

func (n *NodeStatsOSInfo) Json() (string, error) {
	bytes, err := json.Marshal(n)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func (n *NodeStatsOSService) Path() string {
	return node_stats_os_endpoint
}

func (n *NodeStatsOSService) Do(ctx context.Context) (*NodeStatsOSInfo, error) {
	req, err := n.client.newRequest(ctx, "GET", node_stats_os_endpoint, nil)
	if err != nil {
		return nil, err
	}

	res, err := n.client.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	var no NodeStatsOSInfo
	if err := n.client.decodeBody(res, &no); err != nil {
		return nil, err
	}

	return &no, nil
}
