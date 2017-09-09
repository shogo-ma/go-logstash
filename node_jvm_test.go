package logstash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNodeJVMPath(t *testing.T) {
	assert := assert.New(t)

	client := setupTestClient()

	path := client.NodeJVMInfo().Path()
	assert.Equal("/_node/jvm", path)
}
