package yamusic

import (
	"net/http"
	"net/http/httptest"
	"net/url"
)

var (
	// mux is the HTTP request multiplexer used with the test server.
	mux *http.ServeMux

	// client is the Yandex.Music client being tested.
	client *Client

	// server is a test HTTP server used to provide mock API responses.
	server *httptest.Server
)

// setup sets up a test HTTP server along with a yamusic.Client that is
// configured to talk to that test server. Tests should register handlers on
// mux which provide mock responses for the API method being tested.
func setup() {
	// test server
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	url, _ := url.Parse(server.URL + "/")

	// yamusic client configured to use test server
	client = NewClient(BaseURL(url))
}

// teardown closes the test HTTP server.
func teardown() {
	server.Close()
}
