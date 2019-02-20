package client

func (c *Client) Timers() ([]Job, error) {
	var timers []Job
	return timers, c.get("/api/timers", &timers)
}

func (c *Client) Timer(id string) (Job, error) {
	var timer Job
	return timer, c.get("/api/timers/"+id, &timer)
}

func (c *Client) CreateTimer(o Job) error {
	return c.put("/api/timers", &o)
}

func (c *Client) DeleteTimer(id string) error {
	return c.delete("/api/timers/" + id)
}

func (c *Client) UpdateTimer(id string, o Job) error {
	return c.post("/api/timers/"+id, &o)
}
