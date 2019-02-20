package cluster

func (c *client) PhProbes() ([]Probe, error) {
	var probes []Probe
	return probes, c.get("/api/phprobes", &probes)
}

func (c *client) PhProbe(id string) (Probe, error) {
	var outlet Probe
	return outlet, c.get("/api/phprobes/"+id, &outlet)
}

func (c *client) CreatePhProbe(o Probe) error {
	return c.put("/api/phprobes", &o)
}

func (c *client) DeletePhProbe(id string) error {
	return c.delete("/api/phprobes/" + id)
}

func (c *client) UpdatePhProbe(id string, o Probe) error {
	return c.post("/api/phprobes/"+id, &o)
}

func (c *client) CalibratePhProbe(id string, cal CalibrationDetails) error {
	return c.post("/api/phprobes/"+id+"/calibrate", &cal)
}

func (c *client) PhProbeReadings(id string) (StatsResponse, error) {
	var s StatsResponse
	return s, c.get("/api/phprobes/"+id+"/readings", &s)
}
