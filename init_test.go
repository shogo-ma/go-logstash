package logstash

func setupTestClient() *Client {
	client, _ := NewClient("http://localhost:9600")

	return client
}
