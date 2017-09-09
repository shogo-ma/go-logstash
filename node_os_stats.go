package logstash

import (
	"context"
)

type NodeStatsOSInfo struct {
	OS struct {
		CGroup struct {
			CPUAcct struct {
				ControlGroup string `json:"control_group"`
				UsageNanos   int    `json:"usage_nanos"`
			} `json:"cpuacct"`
			CPU struct {
				ControlGroup    string `json:"control_group"`
				CfsPeriosMicros int    `json:"cfs_period_micros"`
				CfsQuotaMicros  int    `json:"cfs_quota_micros"`
				Stat            struct {
					NumberOfElapsedPeriods int `json:"number_of_elapsed_periods"`
					NumberOfTimesThrottled int `json:"number_of_times_throttled"`
					TimeThrottledNanos     int `json:"time_throttled_nanos"`
				} `json:"stat"`
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
