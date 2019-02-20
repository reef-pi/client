package client

func (c *Client) Jacks() ([]Jack, error) {
	var jacks []Jack
	return jacks, c.get("/api/jacks", &jacks)
}

func (c *Client) Jack(id string) (Jack, error) {
	var jack Jack
	return jack, c.get("/api/jacks/"+id, &jack)
}

func (c *Client) CreateJack(o Jack) error {
	return c.put("/api/jacks", &o)
}

func (c *Client) DeleteJack(id string) error {
	return c.delete("/api/jacks/" + id)
}

func (c *Client) UpdateJack(id string, o Jack) error {
	return c.post("/api/jacks/"+id, &o)
}

func (c *Client) ControlJack(id string, p PinValues) error {
	return c.post("/api/jacks/"+id+"/control", &p)
}
