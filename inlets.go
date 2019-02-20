package client

func (c *Client) Inlets() ([]Inlet, error) {
	var inlets []Inlet
	return inlets, c.get("/api/inlets", &inlets)
}

func (c *Client) Inlet(id string) (Inlet, error) {
	var inlet Inlet
	return inlet, c.get("/api/inlets/"+id, &inlet)
}

func (c *Client) CreateInlet(o Inlet) error {
	return c.put("/api/inlets", &o)
}

func (c *Client) DeleteInlet(id string) error {
	return c.delete("/api/inlets/" + id)
}

func (c *Client) UpdateInlet(id string, o Inlet) error {
	return c.post("/api/inlets/"+id, &o)
}

func (c *Client) ReadInlet(id string) (int, error) {
	var i int
	return i, c.postWithResponse("/api/inlets/"+id+"/read", nil, &i)
}
