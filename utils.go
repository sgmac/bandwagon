package bandwagon

import (
	"net/http"
	"net/http/httptest"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
	client *Client
)

func setup() {

	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	creds := Credentials{"empty", "empty"}
	client = NewClient(creds)
	client.BaseURL = server.URL
}
