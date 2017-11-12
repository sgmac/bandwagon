package bandwagon

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
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

func (c *Client) doRequest(apiPath string) (*ServerResponse, error) {

	veid := c.creds.VeID
	apikey := c.creds.APIKey
	baseURL := c.BaseURL

	u := baseURL + apiPath + "veid=" + veid + "&api_key=" + apikey

	ul, _ := url.Parse(u)
	req := &http.Request{
		URL:    ul,
		Method: http.MethodGet,
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response := new(ServerResponse)
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
