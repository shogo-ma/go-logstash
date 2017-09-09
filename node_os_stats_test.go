package logstash

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNodeStatsOSPath(t *testing.T) {
	assert := assert.New(t)
	client := setupTestClient()

	path := client.NodeStatsOSInfo().Path()
	assert.Equal("/_node/stats/os", path)
}

func TestNodeStatsOSDo(t *testing.T) {
	assert := assert.New(t)
	client := setupTestClient()

	res, err := client.NodeStatsOSInfo().Do(context.Background())
	assert.NoError(err)

	j, err := res.Json()
	assert.NoError(err)
	assert.Equal(
		`{"os":{"cgroup":{"cpuacct":{"usage_nanos":56317523679003,"control_group":"/user.slice"},"cpu":{"cfs_quota_micros":-1,"control_group":"/user.slice","stat":{"number_of_times_throttled":0,"time_throttled_nanos":0,"number_of_elapsed_periods":0},"cfs_period_micros":100000}}}}`,
		j,
	)
}
