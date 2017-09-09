package logstash

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNodeJVMPath(t *testing.T) {
	assert := assert.New(t)

	client := setupTestClient()

	path := client.NodeJVMInfo().Path()
	assert.Equal("/_node/jvm", path)
}

func TestNodeJvmDo(t *testing.T) {
	assert := assert.New(t)

	client := setupTestClient()

	res, err := client.NodeJVMInfo().Do(context.Background())
	assert.NoError(err)

	j, err := res.Json()
	assert.NoError(err)
	assert.Equal(
		`{"jvm":{"pid":1,"version":"1.8.0_141","vm_name":"OpenJDK 64-Bit Server VM","vm_version":"1.8.0_141","vm_vendor":"Oracle Corporation","start_time_in_millis":1504967438640,"mem":{"heap_init_in_bytes":1073741824,"heap_max_in_bytes":1038876672,"non_heap_init_in_bytes":2555904,"non_heap_max_in_bytes":0},"gc_collectors":["ParNew","ConcurrentMarkSweep"]}}`,
		j,
	)
}
