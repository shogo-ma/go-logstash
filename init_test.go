package logstash

import (
	"net/http"
	"net/http/httptest"
)

func setupTestClient() *Client {
	mux := http.NewServeMux()
	mux.HandleFunc("/_node/os", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "testdata/node_os.json")
	})

	mux.HandleFunc("/_node/jvm", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "testdata/node_jvm.json")
	})

	mux.HandleFunc("/_node/pipeline", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "testdata/node_pipeline.json")
	})

	mux.HandleFunc("/_node/plugins", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "testdata/node_plugins.json")
	})

	mux.HandleFunc("/_node/hot_threads", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "testdata/hot_threads.json")
	})

	mux.HandleFunc("/_node/stats/os", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "testdata/node_stats_os.json")
	})

	mux.HandleFunc("/_node/stats/pipeline", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "testdata/node_os_pipeline.json")
	})

	mux.HandleFunc("/_node/stats/reloads", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "testdata/node_stats_reloads.json")
	})

	mux.HandleFunc("/_node/stats/process", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "testdata/node_stats_process.json")
	})

	mux.HandleFunc("/_node/stats/jvm", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "testdata/node_stats_jvm.json")
	})

	ts := httptest.NewServer(mux)
	client, _ := NewClient(ts.URL)

	return client
}
