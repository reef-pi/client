package cluster

type AnalogInput struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Pin    int    `json:"pin"`
	Driver string `json:"driver"`
}

func (c *client) AnalogInputs() ([]AnalogInput, error) {
	var ais []AnalogInput
	return ais, c.get("/api/analog_inputs", &ais)
}

func (c *client) AnalogInput(id string) (AnalogInput, error) {
	var ai AnalogInput
	return ai, c.get("/api/analaog_inputs/"+id, &ai)
}

func (c *client) CreateAnalogInput(o AnalogInput) error {
	return c.put("/api/analog_inputs", &o)
}

func (c *client) DeleteAnalogInput(id string) error {
	return c.delete("/api/analog_inputs/" + id)
}

func (c *client) UpdateAnalogInput(id string, o AnalogInput) error {
	return c.post("/api/analog_inputs/"+id, &o)
}
