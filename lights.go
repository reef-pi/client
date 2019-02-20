package client

func (c *Client) Lights() ([]Light, error) {
	var lights []Light
	return lights, c.get("/api/lights", &lights)
}

func (c *Client) Light(id string) (Light, error) {
	var light Light
	return light, c.get("/api/lights/"+id, &light)
}

func (c *Client) CreateLight(o Light) error {
	return c.put("/api/lights", &o)
}

func (c *Client) DeleteLight(id string) error {
	return c.delete("/api/lights/" + id)
}

func (c *Client) UpdateLight(id string, o Light) error {
	return c.post("/api/lights/"+id, &o)
}
