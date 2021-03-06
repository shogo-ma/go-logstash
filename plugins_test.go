package logstash

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPluginsPath(t *testing.T) {
	assert := assert.New(t)

	client := setupTestClient()

	path := client.Plugins().Path()
	assert.Equal("/_node/plugins", path)
}

func TestPluginsDo(t *testing.T) {
	assert := assert.New(t)

	client := setupTestClient()

	res, err := client.Plugins().Do(context.Background())
	assert.NoError(err)

	j, err := res.Json()
	assert.NoError(err)
	assert.Equal(
		`{"total":95,"plugins":[{"name":"logstash-codec-cef","version":"4.1.3"},{"name":"logstash-codec-collectd","version":"3.0.6"},{"name":"logstash-codec-dots","version":"3.0.3"},{"name":"logstash-codec-edn","version":"3.0.3"},{"name":"logstash-codec-edn_lines","version":"3.0.3"},{"name":"logstash-codec-es_bulk","version":"3.0.4"},{"name":"logstash-codec-fluent","version":"3.1.2"},{"name":"logstash-codec-graphite","version":"3.0.3"},{"name":"logstash-codec-json","version":"3.0.3"},{"name":"logstash-codec-json_lines","version":"3.0.3"},{"name":"logstash-codec-line","version":"3.0.3"},{"name":"logstash-codec-msgpack","version":"3.0.3"},{"name":"logstash-codec-multiline","version":"3.0.6"},{"name":"logstash-codec-netflow","version":"3.4.1"},{"name":"logstash-codec-plain","version":"3.0.3"},{"name":"logstash-codec-rubydebug","version":"3.0.3"},{"name":"logstash-filter-clone","version":"3.0.3"},{"name":"logstash-filter-csv","version":"3.0.4"},{"name":"logstash-filter-date","version":"3.1.7"},{"name":"logstash-filter-dissect","version":"1.0.11"},{"name":"logstash-filter-dns","version":"3.0.4"},{"name":"logstash-filter-drop","version":"3.0.3"},{"name":"logstash-filter-fingerprint","version":"3.0.4"},{"name":"logstash-filter-geoip","version":"4.2.1"},{"name":"logstash-filter-grok","version":"3.4.2"},{"name":"logstash-filter-json","version":"3.0.3"},{"name":"logstash-filter-kv","version":"4.0.1"},{"name":"logstash-filter-metrics","version":"4.0.3"},{"name":"logstash-filter-mutate","version":"3.1.5"},{"name":"logstash-filter-ruby","version":"3.0.3"},{"name":"logstash-filter-sleep","version":"3.0.4"},{"name":"logstash-filter-split","version":"3.1.3"},{"name":"logstash-filter-syslog_pri","version":"3.0.3"},{"name":"logstash-filter-throttle","version":"4.0.2"},{"name":"logstash-filter-urldecode","version":"3.0.4"},{"name":"logstash-filter-useragent","version":"3.1.3"},{"name":"logstash-filter-uuid","version":"3.0.3"},{"name":"logstash-filter-xml","version":"4.0.3"},{"name":"logstash-input-beats","version":"3.1.23"},{"name":"logstash-input-couchdb_changes","version":"3.1.2"},{"name":"logstash-input-dead_letter_queue","version":"1.0.5"},{"name":"logstash-input-elasticsearch","version":"4.0.4"},{"name":"logstash-input-exec","version":"3.1.3"},{"name":"logstash-input-file","version":"4.0.2"},{"name":"logstash-input-ganglia","version":"3.1.1"},{"name":"logstash-input-gelf","version":"3.0.5"},{"name":"logstash-input-generator","version":"3.0.3"},{"name":"logstash-input-graphite","version":"3.0.3"},{"name":"logstash-input-heartbeat","version":"3.0.3"},{"name":"logstash-input-http","version":"3.0.5"},{"name":"logstash-input-http_poller","version":"3.3.1"},{"name":"logstash-input-imap","version":"3.0.3"},{"name":"logstash-input-irc","version":"3.0.3"},{"name":"logstash-input-jdbc","version":"4.2.2"},{"name":"logstash-input-kafka","version":"5.1.8"},{"name":"logstash-input-log4j","version":"3.1.0"},{"name":"logstash-input-lumberjack","version":"3.1.2"},{"name":"logstash-input-pipe","version":"3.0.4"},{"name":"logstash-input-rabbitmq","version":"5.2.4"},{"name":"logstash-input-redis","version":"3.1.3"},{"name":"logstash-input-s3","version":"3.1.5"},{"name":"logstash-input-snmptrap","version":"3.0.3"},{"name":"logstash-input-sqs","version":"3.0.4"},{"name":"logstash-input-stdin","version":"3.2.3"},{"name":"logstash-input-syslog","version":"3.2.1"},{"name":"logstash-input-tcp","version":"4.1.2"},{"name":"logstash-input-twitter","version":"3.0.5"},{"name":"logstash-input-udp","version":"3.1.1"},{"name":"logstash-input-unix","version":"3.0.4"},{"name":"logstash-input-xmpp","version":"3.1.4"},{"name":"logstash-output-cloudwatch","version":"3.0.5"},{"name":"logstash-output-csv","version":"3.0.4"},{"name":"logstash-output-elasticsearch","version":"7.3.8"},{"name":"logstash-output-file","version":"4.1.0"},{"name":"logstash-output-graphite","version":"3.1.2"},{"name":"logstash-output-http","version":"4.3.2"},{"name":"logstash-output-irc","version":"3.0.3"},{"name":"logstash-output-kafka","version":"5.1.7"},{"name":"logstash-output-nagios","version":"3.0.3"},{"name":"logstash-output-null","version":"3.0.3"},{"name":"logstash-output-pagerduty","version":"3.0.4"},{"name":"logstash-output-pipe","version":"3.0.3"},{"name":"logstash-output-rabbitmq","version":"4.0.9"},{"name":"logstash-output-redis","version":"3.0.4"},{"name":"logstash-output-s3","version":"4.0.9"},{"name":"logstash-output-sns","version":"4.0.4"},{"name":"logstash-output-sqs","version":"4.0.2"},{"name":"logstash-output-statsd","version":"3.1.2"},{"name":"logstash-output-stdout","version":"3.1.1"},{"name":"logstash-output-tcp","version":"4.0.1"},{"name":"logstash-output-udp","version":"3.0.3"},{"name":"logstash-output-webhdfs","version":"3.0.3"},{"name":"logstash-output-xmpp","version":"3.0.5"},{"name":"logstash-patterns-core","version":"4.1.1"},{"name":"x-pack","version":"5.5.2"}]}`,
		j,
	)
}
