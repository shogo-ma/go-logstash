package logstash

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNodeStatsJVMPath(t *testing.T) {
	assert := assert.New(t)
	client := setupTestClient()

	path := client.NodeJVMStatsInfo().Path()
	assert.Equal("/_node/stats/jvm", path)
}

func TestNodeStatsJVMDo(t *testing.T) {
	assert := assert.New(t)

	client := setupTestClient()
	res, err := client.NodeJVMStatsInfo().Do(context.Background())
	assert.NoError(err)

	j, err := res.Json()
	assert.NoError(err)
	assert.Equal(
		`{"jvm":{"threads":{"count":31,"peak_count":32},"mem":{"heap_used_percent":29,"heap_committed_in_bytes":1038876672,"heap_max_in_bytes":1038876672,"heap_used_in_bytes":310636032,"non_heap_used_in_bytes":91698768,"non_heap_committed_in_bytes":96452608,"pools":{"survivor":{"peak_used_in_bytes":34865152,"used_in_bytes":32575928,"peak_max_in_bytes":34865152,"max_in_bytes":34865152,"committed_in_bytes":34865152},"old":{"peak_used_in_bytes":87393440,"used_in_bytes":36661584,"peak_max_in_bytes":724828160,"max_in_bytes":724828160,"committed_in_bytes":724828160},"young":{"peak_used_in_bytes":279183360,"used_in_bytes":241398520,"peak_max_in_bytes":279183360,"max_in_bytes":279183360,"committed_in_bytes":279183360}}},"gc":{"collectors":{"old":{"collection_time_in_millis":286,"collection_count":2},"young":{"collection_time_in_millis":460,"collection_count":5}}},"uptime_in_millis":93156}}`,
		j,
	)
}
