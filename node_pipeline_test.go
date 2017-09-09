package logstash

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNodePipelinePath(t *testing.T) {
	assert := assert.New(t)

	client := setupTestClient()

	path := client.NodePipelineInfo().Path()
	assert.Equal("/_node/pipeline", path)
}

func TestNodePipelineDo(t *testing.T) {
	assert := assert.New(t)

	client := setupTestClient()

	res, err := client.NodePipelineInfo().Do(context.Background())
	assert.NoError(err)

	j, err := res.Json()
	assert.NoError(err)
	assert.Equal(
		`{"pipeline":{"workers":4,"batch_size":125,"batch_delay":5,"config_reload_automatic":false,"config_reload_interval":3,"id":"main"}}`,
		j,
	)
}
