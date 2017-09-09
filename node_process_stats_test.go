package logstash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNodeProcessStatsPath(t *testing.T) {
	assert := assert.New(t)
	client := setupTestClient()

	path := client.NodeProcessStatsInfo().Path()
	assert.Equal("/_node/stats/process", path)
}
