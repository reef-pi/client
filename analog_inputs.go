package client

func (c *Client) AnalogInputs() ([]AnalogInput, error) {
	var ais []AnalogInput
	return ais, c.get("/api/analog_inputs", &ais)
}

func (c *Client) AnalogInput(id string) (AnalogInput, error) {
	var ai AnalogInput
	return ai, c.get("/api/analaog_inputs/"+id, &ai)
}

func (c *Client) CreateAnalogInput(o AnalogInput) error {
	return c.put("/api/analog_inputs", &o)
}

func (c *Client) DeleteAnalogInput(id string) error {
	return c.delete("/api/analog_inputs/" + id)
}

func (c *Client) UpdateAnalogInput(id string, o AnalogInput) error {
	return c.post("/api/analog_inputs/"+id, &o)
}

func (c *Client) ReadAnalogInput(id string) (float64, error) {
	var f float64
	return f, c.postWithResponse("/api/analog_inputs/"+id+"/read", nil, &f)
}
