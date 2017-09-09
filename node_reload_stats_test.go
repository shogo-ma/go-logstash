package logstash

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNodeReloadStatsPath(t *testing.T) {
	assert := assert.New(t)
	client := setupTestClient()

	path := client.NodeReloadStatsInfo().Path()
	assert.Equal("/_node/stats/reloads", path)
}

func TestNodeReloadStatsDo(t *testing.T) {
	assert := assert.New(t)
	client := setupTestClient()

	res, err := client.NodeReloadStatsInfo().Do(context.Background())
	assert.NoError(err)

	j, err := res.Json()
	assert.NoError(err)
	assert.Equal(
		`{"reloads":{"successes":0,"failures":0}}`,
		j,
	)
}
