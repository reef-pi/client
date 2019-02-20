package client

func (c *Client) SignIn(u, p string) error {
	cred := Credentials{
		User:     u,
		Password: p,
	}
	return c.post("/auth/signin", &cred)
}

func (c *Client) SignOut() error {
	return c.get("/auth/signout", nil)
}

func (c *Client) Capabilities() (Capabilities, error) {
	var caps Capabilities
	return caps, c.get("/api/capabilities", &caps)
}

func (c *Client) PowerOff() error {
	return c.post("/api/admin/poweroff", nil)
}

func (c *Client) Reboot() error {
	return c.post("/api/admin/reboot", nil)
}

func (c *Client) Reload() error {
	return c.post("/api/admin/reload", nil)
}

func (c *Client) Info() (Info, error) {
	var i Info
	return i, c.get("/api/info", &i)
}

func (c *Client) Settings() (Settings, error) {
	var s Settings
	return s, c.get("/api/settings", &s)
}

func (c *Client) UpdateSettings(s Settings) error {
	return c.post("/api/settings", &s)
}

func (c *Client) UpdateCredentials(cr Credentials) error {
	return c.post("/api/credentials", &cr)
}

func (c *Client) Telemetry() (TelemetryConfig, error) {
	var t TelemetryConfig
	return t, c.get("/api/telemetry", &t)
}

func (c *Client) UpdateTelemetry(t TelemetryConfig) error {
	return c.post("/api/telemetry", &t)
}

func (c *Client) SendTestMessage() error {
	return c.post("/api/telemetry/test_message", nil)
}

func (c *Client) ClearErrors() error {
	return c.delete("/api/errors/clear")
}

func (c *Client) DeleteError(id string) error {
	return c.delete("/api/errors/" + id)
}

func (c *Client) Errors() ([]Error, error) {
	var errors []Error
	return errors, c.get("/api/errors", &errors)
}

func (c *Client) Error(id string) (Error, error) {
	var err Error
	return err, c.get("/api/errors/id/"+id, &err)
}

func (c *Client) HealthStats() (StatsResponse, error) {
	var hs StatsResponse
	return hs, c.get("/api/health_stats", &hs)
}

func (c *Client) Dashboard() (Dashboard, error) {
	var ds Dashboard
	return ds, c.get("/api/dashboard", &ds)
}

func (c *Client) UpdateDashboard(ds Dashboard) error {
	return c.post("/api/dashboard", &ds)
}

func (c *Client) DisplayOn() error {
	return c.post("/api/display/on", nil)
}

func (c *Client) DisplayOff() error {
	return c.post("/api/display/off", nil)
}

func (c *Client) Display() (DisplayState, error) {
	var ds DisplayState
	return ds, c.get("/api/display", &ds)
}

func (c *Client) UpdateDisplay(ds DisplayConfig) error {
	return c.post("/api/display", &ds)
}
