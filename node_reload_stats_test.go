package logstash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNodeReloadStatsPath(t *testing.T) {
	assert := assert.New(t)
	client := setupTestClient()

	path := client.NodeReloadStatsInfo().Path()
	assert.Equal("/_node/stats/reloads", path)
}
