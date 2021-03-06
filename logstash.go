package logstash

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"path/filepath"
)

type Client struct {
	URL        *url.URL
	HTTPClient *http.Client
}

func NewClient(host string) (*Client, error) {
	if host == "" {
		return nil, errors.New("Error: host not found.")
	}

	parsedURL, err := url.ParseRequestURI(host)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	return &Client{
		URL:        parsedURL,
		HTTPClient: client,
	}, nil
}

func (c *Client) newRequest(ctx context.Context, method, path string, body io.Reader) (*http.Request, error) {
	u := *c.URL
	u.Path = filepath.Join(c.URL.Path, path)

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	return req, nil
}

func (c *Client) decodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
}

// node info api
func (c *Client) NodePipelineInfo() *NodePipelineService {
	return NewNodePipelineService(c)
}

func (c *Client) NodeOSInfo() *NodeOSService {
	return NewNodeOSService(c)
}

func (c *Client) NodeJVMInfo() *NodeJVMService {
	return NewNodeJVMService(c)
}

// plugins info api
func (c *Client) Plugins() *PluginsService {
	return NewPluginsService(c)
}

// hot threads info api
func (c *Client) HotThreadsInfo() *HotThreadsInfoService {
	return NewHotThreadsInfoService(c)
}

// node stats info api
func (c *Client) NodeStatsOSInfo() *NodeStatsOSService {
	return NewNodeStatsOSService(c)
}

func (c *Client) NodeReloadStatsInfo() *NodeReloadStatsService {
	return NewNodeReloadStatsService(c)
}

func (c *Client) NodeProcessStatsInfo() *NodeProcessStatsService {
	return NewNodeProcessStatsService(c)
}

func (c *Client) NodeJVMStatsInfo() *NodeJVMStatsService {
	return NewNodeJVMStatsService(c)
}

func (c *Client) NodePipelineStatsInfo() *NodePipelineStatsService {
	return NewNodePipelineStatsService(c)
}
