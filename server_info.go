package bandwagon

import (
	"encoding/json"
	"net/http"
	"net/url"
)

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
