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
	SSHPort             int16
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

// Hostname sets a new hostname for vps.
func (c *Client) Hostname(host string) (*ServerResponse, error) {
	apiPath := fmt.Sprintf("/v1/setHostname?newHostname=%s&", host)
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

// InfoVPS holds information regarding a VPS.
type InfoVPS struct {
	VMType                string            `json:"vm_type"`
	Hostname              string            `json:"hostname"`
	NodeIP                string            `json:"node_ip"`
	NodeAlias             string            `json:"node_alias"`
	NodeLocation          string            `json:"node_location"`
	LocationIPv6Ready     bool              `json:"location_ipv6_ready"`
	Plan                  string            `json:"plan"`
	PlanMonthlyData       int64             `json:"plan_monthly_data"`
	MonthlyDataMultiplier int64             `json:"plan_monthly_data"`
	PlanDisk              int64             `json:"plan_disk"`
	PlanRAM               int32             `json:"plan_ram"`
	PlanSwap              int32             `json:"plan_swap"`
	PlanMaxIPv6           int32             `json:"plan_max_i_pv_6"`
	OS                    string            `json:"os"`
	Email                 string            `json:"email"`
	DataCounter           int32             `json:"data_counter"`
	DataNextReset         int32             `json:"data_next_reset"`
	IPAddresses           []string          `json:"ip_addresses"`
	RDNSApiAvailable      bool              `json:"rdns_api_available"`
	PTR                   map[string]string `json:"ptr"`
	Suspended             bool              `json:"suspended"`
	Error                 int32             `json:"error"`
}

// Info returns details about a VPS.
func (c *Client) Info() (*InfoVPS, error) {
	veid := c.creds.VeID
	apikey := c.creds.APIKey
	apiPath := "/v1/getServiceInfo"
	baseURL := c.BaseURL

	u := baseURL + apiPath + "?veid=" + veid + "&api_key=" + apikey

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

	response := new(InfoVPS)
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
