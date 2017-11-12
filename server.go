package bandwagon

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
