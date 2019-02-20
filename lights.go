package client

func (c *client) Lights() ([]Light, error) {
	var lights []Light
	return lights, c.get("/api/lights", &lights)
}

func (c *client) Light(id string) (Light, error) {
	var light Light
	return light, c.get("/api/lights/"+id, &light)
}

func (c *client) CreateLight(o Light) error {
	return c.put("/api/lights", &o)
}

func (c *client) DeleteLight(id string) error {
	return c.delete("/api/lights/" + id)
}

func (c *client) UpdateLight(id string, o Light) error {
	return c.post("/api/lights/"+id, &o)
}
