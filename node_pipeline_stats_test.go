package logstash

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNodePipelineStatsPath(t *testing.T) {
	assert := assert.New(t)
	client := setupTestClient()

	path := client.NodePipelineStatsInfo().Path()
	assert.Equal("/_node/stats/pipeline", path)
}

func TestNodePipelineStatsDo(t *testing.T) {
	assert := assert.New(t)

	client := setupTestClient()

	res, err := client.NodePipelineStatsInfo().Do(context.Background())
	assert.NoError(err)

	j, err := res.Json()
	assert.NoError(err)
	assert.Equal(
		`{"pipeline":{"events":{"duration_in_millis":0,"in":0,"filtered":0,"out":0,"queue_push_duration_in_millis":0},"plugins":{"inputs":[{"id":"a35197a509596954e905e38521bae12b1498b17d-1","events":{"out":0,"queue_push_duration_in_millis":0},"name":"beats"}],"filters":[],"outputs":[{"id":"a35197a509596954e905e38521bae12b1498b17d-2","name":"stdout"}]},"reloads":{"last_error":null,"successes":0,"last_success_timestamp":null,"last_failure_timestamp":null,"failures":0},"queue":{"type":"memory"},"id":"main"}}`,
		j,
	)
}
