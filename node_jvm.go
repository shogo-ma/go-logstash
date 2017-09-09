package logstash

import (
	"context"
	"encoding/json"
)

type NodeInfoJVM struct {
	JVM struct {
		PID               int    `json:"pid"`
		Version           string `json:"version"`
		VMName            string `json:"vm_name"`
		VMVersion         string `json:"vm_version"`
		VMVendor          string `json:"vm_vendor"`
		StartTimeInMillis int    `json:"start_time_in_millis"`
		Mem               struct {
			HeapInitInBytes    int `json:"heap_init_in_bytes"`
			HeapMaxInBytes     int `json:"heap_max_in_bytes"`
			NonHeapInitInBytes int `json:"non_heap_init_in_bytes"`
			NonHeapMaxInBytes  int `json:"non_heap_max_in_bytes"`
		} `json:"mem"`
		GcCollectors []string `json:"gc_collectors"`
	} `json:"jvm"`
}

type NodeJVMService struct {
	client *Client
}

const (
	node_jvm_endpoint = "/_node/jvm"
)

func NewNodeJVMService(client *Client) *NodeJVMService {
	return &NodeJVMService{client: client}
}

func (n *NodeInfoJVM) Json() (string, error) {
	bytes, err := json.Marshal(n)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func (n *NodeJVMService) Path() string {
	return node_jvm_endpoint
}

func (n *NodeJVMService) Do(ctx context.Context) (*NodeInfoJVM, error) {
	req, err := n.client.newRequest(ctx, "GET", node_jvm_endpoint, nil)
	if err != nil {
		return nil, err
	}

	res, err := n.client.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	var nj NodeInfoJVM
	if err := n.client.decodeBody(res, &nj); err != nil {
		return nil, err
	}

	return &nj, nil
}
