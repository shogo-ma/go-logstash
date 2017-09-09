package logstash

import (
	"context"
	"time"
)

type NodePipelineStatsInfo struct {
	Pipeline struct {
		Events struct {
			DurationInMillis          int `json:"duration_in_millis"`
			In                        int `json:"in"`
			Filtered                  int `json:"filtered"`
			Out                       int `json:"out"`
			QueuePushDurationInMillis int `json:"queue_push_duration_in_millis"`
		} `json:"events"`
		Plugins struct {
			Inputs  []PluginInput   `json:"inputs"`
			Filters []PluginFilters `json:"filters"`
			Outputs []PluginOutputs `json:"outputs"`
		} `json:"plugins"`
		Reloads struct {
			LastError            *time.Time `json:"last_error"`
			Successes            int        `json:"successes"`
			LastSuccessTimestamp *time.Time `json:"last_success_timestamp"`
			LastFailureTimestamp *time.Time `json:"last_success_timestamp"`
			Failures             int        `json:"failures"`
		} `json:"reloads"`
		Queue struct {
			Events   int    `json:"events"`
			Type     string `json:"type"`
			Capacity struct {
				PageCapacityInBytes int `json:"page_capacity_in_bytes"`
				MaxQueueSizeInBytes int `json:"max_queue_size_in_bytes"`
				MaxUnreadEvents     int `json:"max_unread_events"`
			} `json:"capacity"`
			Data struct {
				Path             string `json:"path"`
				FreeSpaceInBytes int    `json:"free_space_in_bytes"`
				StorageType      string `json:"storage_type"`
			} `json:"data"`
		} `json:"queue"`
		ID string `json:"id"`
	} `json:"pipeline"`
}

type PluginInput struct {
	ID     string `json:"id"`
	Events struct {
		Out                       int "out"
		QueuePushDurationInMillis int `json:"queue_push_duration_in_millis"`
	} `json:"events"`
	Name string `json:"name"`
}

type PluginFilters struct {
	ID     string `json:"id"`
	Events struct {
		In                        int "in"
		Out                       int "out"
		QueuePushDurationInMillis int `json:"queue_push_duration_in_millis"`
	} `json:"events"`
	Matches          int `json:"matches"`
	PatternsPerField struct {
		Message string `json:"message"`
	} `json:"patterns_per_field"`
	Name string `json:"name"`
}

type PluginOutputs struct {
	ID     string `json:"id"`
	Events struct {
		In                        int "in"
		Out                       int "out"
		QueuePushDurationInMillis int `json:"queue_push_duration_in_millis"`
	} `json:"events"`
	Name string `json:"name"`
}

type NodePipelineStatsService struct {
	client *Client
}

const (
	node_pipeline_stats_endpoint = "/_node/stats/pipeline"
)

func NewNodePipelineStatsService(client *Client) *NodePipelineStatsService {
	return &NodePipelineStatsService{client: client}
}

func (n *NodePipelineStatsService) Path() string {
	return node_pipeline_stats_endpoint
}

func (n *NodePipelineStatsService) Do(ctx context.Context) (*NodePipelineStatsInfo, error) {
	req, err := n.client.newRequest(ctx, "GET", node_pipeline_stats_endpoint, nil)
	if err != nil {
		return nil, err
	}

	res, err := n.client.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	var nr NodePipelineStatsInfo
	if err := n.client.decodeBody(res, &nr); err != nil {
		return nil, err
	}

	return &nr, nil
}
