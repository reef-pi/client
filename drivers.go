package client

func (c *Client) Drivers() ([]Driver, error) {
	var ds []Driver
	return ds, c.get("/api/drivers", &ds)
}

func (c *Client) Driver(id string) (Driver, error) {
	var d Driver
	return d, c.get("/api/drivers/"+id, &d)
}

func (c *Client) CreateDriver(o Driver) error {
	return c.put("/api/drivers", &o)
}

func (c *Client) DeleteDriver(id string) error {
	return c.delete("/api/drivers/" + id)
}

func (c *Client) UpdateDriver(id string, o Driver) error {
	return c.post("/api/drivers/"+id, &o)
}
