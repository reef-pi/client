package client

func (c *Client) PhProbes() ([]Probe, error) {
	var probes []Probe
	return probes, c.get("/api/phprobes", &probes)
}

func (c *Client) PhProbe(id string) (Probe, error) {
	var outlet Probe
	return outlet, c.get("/api/phprobes/"+id, &outlet)
}

func (c *Client) CreatePhProbe(o Probe) error {
	return c.put("/api/phprobes", &o)
}

func (c *Client) DeletePhProbe(id string) error {
	return c.delete("/api/phprobes/" + id)
}

func (c *Client) UpdatePhProbe(id string, o Probe) error {
	return c.post("/api/phprobes/"+id, &o)
}

func (c *Client) CalibratePhProbe(id string, cal CalibrationDetails) error {
	return c.post("/api/phprobes/"+id+"/calibrate", &cal)
}

func (c *Client) PhProbeReadings(id string) (StatsResponse, error) {
	var s StatsResponse
	return s, c.get("/api/phprobes/"+id+"/readings", &s)
}
