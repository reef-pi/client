package client

func (c *Client) ATOs() ([]ATO, error) {
	var atos []ATO
	return atos, c.get("/api/atos", &atos)
}

func (c *Client) ATO(id string) (ATO, error) {
	var ato ATO
	return ato, c.get("/api/atos/"+id, &ato)
}

func (c *Client) CreateATO(o ATO) error {
	return c.put("/api/atos", &o)
}

func (c *Client) DeleteATO(id string) error {
	return c.delete("/api/atos/" + id)
}

func (c *Client) UpdateATO(id string, o ATO) error {
	return c.post("/api/atos/"+id, &o)
}

func (c *Client) ATOUsage(id string) (StatsResponse, error) {
	var s StatsResponse
	return s, c.get("/api/atos/"+id+"/usage", &s)
}
