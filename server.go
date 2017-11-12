package bandwagon

import (
	"encoding/json"
	"fmt"
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

// InstallResposne holds the response for installing an OS.
type InstallResponse struct {
	Error               int32
	Message             string
	AdditionalErrorCode int64
	AdditionalErrorInfo string
	RootPassword        string
	SSHPort             string
	NotificationEmail   string
}

// Reboot restarts the vps with a given VeID
func (c *Client) Reboot() (*ServerResponse, error) {
	apiPath := "/v1/restart?"
	response, err := c.doRequest(apiPath)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Stop stops a vps.
func (c *Client) Stop() (*ServerResponse, error) {
	apiPath := "/v1/stop?"
	response, err := c.doRequest(apiPath)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Start starts a vps.
func (c *Client) Start() (*ServerResponse, error) {
	apiPath := "/v1/start?"
	response, err := c.doRequest(apiPath)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Kill kills a vps.
func (c *Client) Kill() (*ServerResponse, error) {
	apiPath := "/v1/kill?"
	response, err := c.doRequest(apiPath)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Install installs a given OS on a VPS.
func (c *Client) Install(os string) (*InstallResponse, error) {
	apiPath := fmt.Sprintf("/v1/reinstallOS?os=%s&", os)
	veid := c.creds.VeID
	apikey := c.creds.APIKey
	baseURL := c.BaseURL

	u := baseURL + apiPath + "veid=" + veid + "&api_key=" + apikey

	fmt.Println(u)
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

	response := new(InstallResponse)
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
