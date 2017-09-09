package logstash

import (
	"context"
	"encoding/json"
)

type NodeJVMStatsInfo struct {
	JVM struct {
		Threads struct {
			Count     int `json:"count"`
			PeakCount int `json:"peak_count"`
		} `json:"threads"`
		Mem struct {
			HeapUsedPercent         int `json:"heap_used_percent"`
			HeapCommittedInBytes    int `json:"heap_committed_in_bytes"`
			HeapMaxInBytes          int `json:"heap_max_in_bytes"`
			HeapUsedInBytes         int `json:"heap_used_in_bytes"`
			NonHeapUsedInBytes      int `json:"non_heap_used_in_bytes"`
			NonHeapCommittedInBytes int `json:"non_heap_committed_in_bytes"`
			Pools                   struct {
				Survivor struct {
					PeakUsedInBytes  int `json:"peak_used_in_bytes"`
					UsedInBytes      int `json:"used_in_bytes"`
					PeakMaxInBytes   int `json:"peak_max_in_bytes"`
					MaxInBytes       int `json:"max_in_bytes"`
					CommittedInBytes int `json:"committed_in_bytes"`
				} `json:"survivor"`
				Old struct {
					PeakUsedInBytes  int `json:"peak_used_in_bytes"`
					UsedInBytes      int `json:"used_in_bytes"`
					PeakMaxInBytes   int `json:"peak_max_in_bytes"`
					MaxInBytes       int `json:"max_in_bytes"`
					CommittedInBytes int `json:"committed_in_bytes"`
				} `json:"old"`
				Young struct {
					PeakUsedInBytes  int `json:"peak_used_in_bytes"`
					UsedInBytes      int `json:"used_in_bytes"`
					PeakMaxInBytes   int `json:"peak_max_in_bytes"`
					MaxInBytes       int `json:"max_in_bytes"`
					CommittedInBytes int `json:"committed_in_bytes"`
				} `json:"young"`
			} `json:"pools"`
		} `json:"mem"`
		GC struct {
			Collectors struct {
				Old struct {
					CollectionTimeInMillis int `json:"collection_time_in_millis"`
					CollectionCount        int `json:"collection_count"`
				} `json:"old"`
				Young struct {
					CollectionTimeInMillis int `json:"collection_time_in_millis"`
					CollectionCount        int `json:"collection_count"`
				} `json:"young"`
			} `json:"collectors"`
		} `json:"gc"`
		UptimeInMillis int `json:"uptime_in_millis"`
	} `json:"jvm"`
}

type NodeJVMStatsService struct {
	client *Client
}

const (
	node_jvm_stats_endpoint = "/_node/stats/jvm"
)

func NewNodeJVMStatsService(client *Client) *NodeJVMStatsService {
	return &NodeJVMStatsService{client: client}
}

func (n *NodeJVMStatsService) Path() string {
	return node_jvm_stats_endpoint
}

func (n *NodeJVMStatsInfo) Json() (string, error) {
	bytes, err := json.Marshal(n)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func (n *NodeJVMStatsService) Do(ctx context.Context) (*NodeJVMStatsInfo, error) {
	req, err := n.client.newRequest(ctx, "GET", node_jvm_stats_endpoint, nil)
	if err != nil {
		return nil, err
	}

	res, err := n.client.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	var nr NodeJVMStatsInfo
	if err := n.client.decodeBody(res, &nr); err != nil {
		return nil, err
	}

	return &nr, nil
}
