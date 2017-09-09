package logstash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHotThreadsInfoPath(t *testing.T) {
	assert := assert.New(t)
	client := setupTestClient()

	path := client.HotThreadsInfo().Path()
	assert.Equal("/_node/hot_threads", path)
}
