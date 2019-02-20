package client

func (c *client) Drivers() ([]Driver, error) {
	var ds []Driver
	return ds, c.get("/api/drivers", &ds)
}

func (c *client) Driver(id string) (Driver, error) {
	var d Driver
	return d, c.get("/api/drivers/"+id, &d)
}

func (c *client) CreateDriver(o Driver) error {
	return c.put("/api/drivers", &o)
}

func (c *client) DeleteDriver(id string) error {
	return c.delete("/api/drivers/" + id)
}

func (c *client) UpdateDriver(id string, o Driver) error {
	return c.post("/api/drivers/"+id, &o)
}
