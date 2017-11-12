package bandwagon

import (
	"encoding/json"
	"net/http"
	"net/url"
)

// Server provides basic operations for a given vps.
type Server interface {
	Reboot() (*ServerResponse, error)
}

// ServerResponse holds the Error and Message in case of failure.
// Error:0 means no error.
type ServerResponse struct {
	Error   int32  `json:"error"`
	Message string `json:"message,omitempty"`
}

// Reboot restarts the vps with a given VeID
func (c *Client) Reboot() (*ServerResponse, error) {

	veid := c.creds.VeID
	apikey := c.creds.APIKey
	apiPath := "/v1/restart?"
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
