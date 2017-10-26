// Package bandwagon implements BandwagonHOST API.
package bandwagon

import (
	"net/http"
)

const (
	defaultbaseURL = "https://api.64clouds.com/"
)

// Client communicates with BandwagonHOST.
type Client struct {
	// client performs HTTP calls.
	http *http.Client

	// BaseURL for the API.
	BaseURL string

	// Credentials used for http connections.
	creds Credentials

	// VirtualServer manage basic operations with a VPS.
	VirtualServer
}

// Credentials keeps the api_key and the virtual eid.
type Credentials struct {
	APIKey string
	VeID   string
}

// NewClient returns a new client.
func NewClient(cred Credentials) *Client {
	c := &Client{
		http:    http.DefaultClient,
		BaseURL: defaultbaseURL,
		creds:   cred,
	}
	return c
}
