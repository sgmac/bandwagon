package bandwagon

import (
	"encoding/json"
	"net/http"
	"net/url"
)

// Images holds the JSON response for the API call.
type Images struct {
	Templates []string `json:"templates"`
	Installed string   `json:"installed"`
	Error     int32    `json:"error"`
}

// ListImages returns the available OS images.
func (c *Client) ListImages() (*Images, error) {

	veid := c.creds.VeID
	apikey := c.creds.APIKey
	apiPath := "/v1" + "/getAvailableOS?"
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

	os := new(Images)

	err = json.NewDecoder(resp.Body).Decode(&os)
	if err != nil {
		return nil, err
	}
	return os, nil
}
