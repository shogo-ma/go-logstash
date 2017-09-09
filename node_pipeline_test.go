package logstash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNodePipelinePath(t *testing.T) {
	assert := assert.New(t)

	client := setupTestClient()

	path := client.NodePipelineInfo().Path()
	assert.Equal("/_node/pipeline", path)
}
