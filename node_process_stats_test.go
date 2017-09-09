package logstash

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNodeProcessStatsPath(t *testing.T) {
	assert := assert.New(t)
	client := setupTestClient()

	path := client.NodeProcessStatsInfo().Path()
	assert.Equal("/_node/stats/process", path)
}

func TestNodeProcessStatsDo(t *testing.T) {
	assert := assert.New(t)
	client := setupTestClient()

	res, err := client.NodeProcessStatsInfo().Do(context.Background())
	assert.NoError(err)

	j, err := res.Json()
	assert.NoError(err)
	assert.Equal(
		`{"process":{"open_file_descriptors":109,"peak_open_file_descriptors":109,"max_file_descriptors":1048576,"mem":{"total_virtual_in_bytes":4855468032},"cpu":{"total_in_millis":63750,"percent":3,"load_average":{"1m":2.05,"5m":2.16,"15m":2.13}}}}`,
		j,
	)
}
