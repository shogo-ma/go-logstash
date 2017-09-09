package logstash

import (
	"context"
)

type NodeInfoPipeline struct {
	Pipeline struct {
		Workers               int    `json:"workers"`
		BatchSize             int    `json:"batch_size"`
		BatchDelay            int    `json:"batch_delay"`
		ConfigReloadAutomatic bool   `json:"config_reload_automatic"`
		ConfigReloadInterval  int    `json:"config_reload_interval"`
		ID                    string `json:"id"`
	}
}

type NodePipelineService struct {
	client *Client
}

const (
	node_pipeline_endpoint = "/_node/pipeline"
)

func NewNodePipelineService(client *Client) *NodePipelineService {
	return &NodePipelineService{client: client}
}

func (n *NodePipelineService) Path() string {
	return node_pipeline_endpoint
}

func (n *NodePipelineService) Do(ctx context.Context) (*NodeInfoPipeline, error) {
	req, err := n.client.newRequest(ctx, "GET", node_pipeline_endpoint, nil)
	if err != nil {
		return nil, err
	}

	res, err := n.client.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	var np NodeInfoPipeline
	if err := n.client.decodeBody(res, &np); err != nil {
		return nil, err
	}

	return &np, nil
}
