package logstash

import "context"

type NodeProcessStatsInfo struct {
	Process struct {
		OpenFileDescriptors     int `json:"open_file_descriptors"`
		PeakOpenFileDescriptors int `json:"peak_open_file_descriptors"`
		MaxFileDescriptors      int `json:"max_file_descriptors"`
		Mem                     struct {
			TotalVirtualInBytes int `json:"total_virtual_in_bytes"`
		} `json:"mem"`
		CPU struct {
			TotalInMillis int `json:"total_in_millis"`
			Percent       int `json:"percent"`
			LoadAverage   struct {
				Minitue float64 `json:"1m"`
			} `json:"load_average"`
		} `json:"cpu"`
	} `json:"process"`
}

type NodeProcessStatsService struct {
	client *Client
}

const (
	node_process_stats_endpoint = "/_node/stats/process"
)

func NewNodeProcessStatsService(client *Client) *NodeProcessStatsService {
	return &NodeProcessStatsService{client: client}
}

func (n *NodeProcessStatsService) Path() string {
	return node_process_stats_endpoint
}

func (n *NodeProcessStatsService) Do(ctx context.Context) (*NodeProcessStatsInfo, error) {
	req, err := n.client.newRequest(ctx, "GET", node_process_stats_endpoint, nil)
	if err != nil {
		return nil, err
	}

	res, err := n.client.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	var np NodeProcessStatsInfo
	if err := n.client.decodeBody(res, &np); err != nil {
		return nil, err
	}

	return &np, nil
}
