package cluster

func (c *client) SignIn(u, p string) error {
	cred := Credentials{
		User:     u,
		Password: p,
	}
	return c.post("/auth/signin", &cred)
}

func (c *client) SignOut() error {
	return c.get("/auth/signout", nil)
}

func (c *client) Capabilities() (Capabilities, error) {
	var caps Capabilities
	return caps, c.get("/api/capabilities", &caps)
}

func (c *client) PowerOff() error {
	return c.post("/api/admin/poweroff", nil)
}

func (c *client) Reboot() error {
	return c.post("/api/admin/reboot", nil)
}

func (c *client) Reload() error {
	return c.post("/api/admin/reload", nil)
}

func (c *client) Info() (Info, error) {
	var i Info
	return i, c.get("/api/info", &i)
}

func (c *client) Settings() (Settings, error) {
	var s Settings
	return s, c.get("/api/settings", &s)
}

func (c *client) UpdateSettings(s Settings) error {
	return c.post("/api/settings", &s)
}

func (c *client) UpdateCredentials(cr Credentials) error {
	return c.post("/api/credentials", &cr)
}

func (c *client) Telemetry() (TelemetryConfig, error) {
	var t TelemetryConfig
	return t, c.get("/api/telemetry", &t)
}

func (c *client) UpdateTelemetry(t TelemetryConfig) error {
	return c.post("/api/telemetry", &t)
}

func (c *client) SendTestMessage() error {
	return c.post("/api/telemetry/test_message", nil)
}

func (c *client) ClearErrors() error {
	return c.delete("/api/errors/clear")
}

func (c *client) DeleteError(id string) error {
	return c.delete("/api/errors/" + id)
}

func (c *client) Errors() ([]Error, error) {
	var errors []Error
	return errors, c.get("/api/errors", &errors)
}

func (c *client) Error(id string) (Error, error) {
	var err Error
	return err, c.get("/api/errors/id/"+id, &err)
}

func (c *client) HealthStats() (StatsResponse, error) {
	var hs StatsResponse
	return hs, c.get("/api/health_stats", &hs)
}

func (c *client) Dashboard() (Dashboard, error) {
	var ds Dashboard
	return ds, c.get("/api/dashboard", &ds)
}

func (c *client) UpdateDashboard(ds Dashboard) error {
	return c.post("/api/dashboard", &ds)
}

func (c *client) DisplayOn() error {
	return c.post("/api/display/on", nil)
}

func (c *client) DisplayOff() error {
	return c.post("/api/display/off", nil)
}

func (c *client) Display() (DisplayState, error) {
	var ds DisplayState
	return ds, c.get("/api/display", &ds)
}

func (c *client) UpdateDisplay(ds DisplayConfig) error {
	return c.post("/api/display", &ds)
}
