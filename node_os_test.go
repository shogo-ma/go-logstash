package logstash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNodeOSPath(t *testing.T) {
	assert := assert.New(t)

	client := setupTestClient()

	path := client.NodeOSInfo().Path()
	assert.Equal("/_node/os", path)
}
