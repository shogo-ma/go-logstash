package logstash

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHotThreadsInfoPath(t *testing.T) {
	assert := assert.New(t)
	client := setupTestClient()

	path := client.HotThreadsInfo().Path()
	assert.Equal("/_node/hot_threads", path)
}

func TestHotThreadDo(t *testing.T) {
	assert := assert.New(t)

	client := setupTestClient()

	res, err := client.HotThreadsInfo().Do(context.Background())
	assert.NoError(err)

	j, err := res.Json()
	assert.NoError(err)
	assert.Equal(
		`{"hot_threads":{"time":"2017-09-09T14:34:02Z","busiest_threads":3,"threads":[{"name":"LogStash::Runner","thread_id":1,"percent_of_cpu_time":5.27,"state":"timed_waiting","traces":["java.lang.Object.wait(Native Method)","java.lang.Thread.join(Thread.java:1260)","org.jruby.internal.runtime.NativeThread.join(NativeThread.java:75)"]},{"name":"[main]\u003eworker3","thread_id":27,"percent_of_cpu_time":2.18,"state":"waiting","traces":["sun.misc.Unsafe.park(Native Method)","java.util.concurrent.locks.LockSupport.park(LockSupport.java:175)","java.util.concurrent.locks.AbstractQueuedSynchronizer.parkAndCheckInterrupt(AbstractQueuedSynchronizer.java:836)"]},{"name":"[main]\u003eworker1","thread_id":25,"percent_of_cpu_time":2.16,"state":"waiting","traces":["sun.misc.Unsafe.park(Native Method)","java.util.concurrent.locks.LockSupport.park(LockSupport.java:175)","java.util.concurrent.locks.AbstractQueuedSynchronizer.parkAndCheckInterrupt(AbstractQueuedSynchronizer.java:836)"]}]}}`,
		j,
	)
}
