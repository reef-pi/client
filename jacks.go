package client

func (c *client) Jacks() ([]Jack, error) {
	var jacks []Jack
	return jacks, c.get("/api/jacks", &jacks)
}

func (c *client) Jack(id string) (Jack, error) {
	var jack Jack
	return jack, c.get("/api/jacks/"+id, &jack)
}

func (c *client) CreateJack(o Jack) error {
	return c.put("/api/jacks", &o)
}

func (c *client) DeleteJack(id string) error {
	return c.delete("/api/jacks/" + id)
}

func (c *client) UpdateJack(id string, o Jack) error {
	return c.post("/api/jacks/"+id, &o)
}

func (c *client) ControlJack(id string, p PinValues) error {
	return c.post("/api/jacks/"+id+"/control", &p)
}
