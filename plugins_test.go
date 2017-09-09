package logstash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPluginsPath(t *testing.T) {
	assert := assert.New(t)

	client := setupTestClient()

	path := client.Plugins().Path()
	assert.Equal("/_node/plugins", path)
}
