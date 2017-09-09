package logstash

import (
	"context"
)

type NodeInfoOS struct {
	Name                string `json:"name"`
	Arch                string `json:"arch"`
	Version             string `json:"version"`
	AvailableProcessors int    `json:"available_processors"`
}

type NodeOSService struct {
	client *Client
}

const (
	node_os_endpoint = "/_node/os"
)

func NewNodeOSService(client *Client) *NodeOSService {
	return &NodeOSService{client: client}
}

func (n *NodeOSService) Path() string {
	return node_os_endpoint
}

func (n *NodeOSService) Do(ctx context.Context) (*NodeInfoOS, error) {
	req, err := n.client.newRequest(ctx, "GET", node_os_endpoint, nil)
	if err != nil {
		return nil, err
	}

	res, err := n.client.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	var no NodeInfoOS
	if err := n.client.decodeBody(res, &no); err != nil {
		return nil, err
	}

	return &no, nil
}
