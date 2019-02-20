package client

func (c *Client) Outlets() ([]Outlet, error) {
	var outlets []Outlet
	return outlets, c.get("/api/outlets", &outlets)
}

func (c *Client) Outlet(id string) (Outlet, error) {
	var outlet Outlet
	return outlet, c.get("/api/outlets/"+id, &outlet)
}

func (c *Client) CreateOutlet(o Outlet) error {
	return c.put("/api/outlets", &o)
}

func (c *Client) DeleteOutlet(id string) error {
	return c.delete("/api/outlets/" + id)
}

func (c *Client) UpdateOutlet(id string, o Outlet) error {
	return c.post("/api/outlets/"+id, &o)
}
