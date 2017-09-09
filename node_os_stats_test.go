package logstash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNodeStatsOSPath(t *testing.T) {
	assert := assert.New(t)
	client := setupTestClient()

	path := client.NodeStatsOSInfo().Path()
	assert.Equal("/_node/stats/os", path)
}
