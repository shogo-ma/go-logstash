package logstash

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNodeOSPath(t *testing.T) {
	assert := assert.New(t)

	client := setupTestClient()

	path := client.NodeOSInfo().Path()
	assert.Equal("/_node/os", path)
}

func TestNodeDo(t *testing.T) {
	assert := assert.New(t)

	client := setupTestClient()

	res, err := client.NodeOSInfo().Do(context.Background())
	assert.NoError(err)

	j, err := res.Json()
	assert.NoError(err)
	assert.Equal(
		`{"os":{"name":"Linux","arch":"amd64","version":"4.4.0-63-generic","available_processors":4}}`,
		j,
	)
}
