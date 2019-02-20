package cluster

func (c *client) Dosers() ([]Pump, error) {
	var dosers []Pump
	return dosers, c.get("/api/doser/pumps", &dosers)
}

func (c *client) Doser(id string) (Pump, error) {
	var doser Pump
	return doser, c.get("/api/doser/pumps/"+id, &doser)
}

func (c *client) CreateDoser(o Pump) error {
	return c.put("/api/doser/pumps", &o)
}

func (c *client) DeleteDoser(id string) error {
	return c.delete("/api/doser/pumps/" + id)
}

func (c *client) UpdateDoser(id string, o Pump) error {
	return c.post("/api/doser/pumps/"+id, &o)
}

func (c *client) DoserUsage(id string) (StatsResponse, error) {
	var s StatsResponse
	return s, c.get("/api/doser/pumps/"+id+"/usage", &s)
}

func (c *client) CalibrateDoser(id string, cal CalibrationDetails) error {
	return c.post("/api/doser/pumps/"+id+"/calibrate", &cal)
}

func (c *client) ScheduleDoser(id string, s DosingRegiment) error {
	return c.post("/api/doser/pumps/"+id+"/schedule", &s)
}
