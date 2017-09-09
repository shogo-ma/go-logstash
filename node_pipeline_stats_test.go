package logstash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNodePipelineStatsPath(t *testing.T) {
	assert := assert.New(t)
	client := setupTestClient()

	path := client.NodePipelineStatsInfo().Path()
	assert.Equal("/_node/stats/pipeline", path)
}
