package logstash

import (
	"context"
)

type Plugins struct {
	Total   int      `json:"total"`
	Plugins []Plugin `json:"plugins"`
}

type Plugin struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type PluginsService struct {
	client *Client
}

const (
	plugins_endpoint = "/_node/plugins"
)

func NewPluginsService(client *Client) *PluginsService {
	return &PluginsService{client: client}
}

func (p *PluginsService) Path() string {
	return plugins_endpoint
}

func (p *PluginsService) Do(ctx context.Context) (*Plugins, error) {
	req, err := p.client.newRequest(ctx, "GET", plugins_endpoint, nil)
	if err != nil {
		return nil, err
	}

	res, err := p.client.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	var ps Plugins
	if err := p.client.decodeBody(res, &ps); err != nil {
		return nil, err
	}

	return &ps, nil
}
