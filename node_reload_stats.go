package logstash

import (
	"context"
	"encoding/json"
)

type NodeReloadStatsInfo struct {
	Reloads struct {
		Successes int `json:"successes"`
		Failures  int `json:"failures"`
	} `json:"reloads"`
}

type NodeReloadStatsService struct {
	client *Client
}

const (
	node_reload_stats_endpoint = "/_node/stats/reloads"
)

func NewNodeReloadStatsService(client *Client) *NodeReloadStatsService {
	return &NodeReloadStatsService{client: client}
}

func (n *NodeReloadStatsInfo) Json() (string, error) {
	bytes, err := json.Marshal(n)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func (n *NodeReloadStatsService) Path() string {
	return node_reload_stats_endpoint
}

func (n *NodeReloadStatsService) Do(ctx context.Context) (*NodeReloadStatsInfo, error) {
	req, err := n.client.newRequest(ctx, "GET", node_reload_stats_endpoint, nil)
	if err != nil {
		return nil, err
	}

	res, err := n.client.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	var nr NodeReloadStatsInfo
	if err := n.client.decodeBody(res, &nr); err != nil {
		return nil, err
	}

	return &nr, nil
}
