package cluster

func (c *client) Timers() ([]Job, error) {
	var timers []Job
	return timers, c.get("/api/timers", &timers)
}

func (c *client) Timer(id string) (Job, error) {
	var timer Job
	return timer, c.get("/api/timers/"+id, &timer)
}

func (c *client) CreateTimer(o Job) error {
	return c.put("/api/timers", &o)
}

func (c *client) DeleteTimer(id string) error {
	return c.delete("/api/timers/" + id)
}

func (c *client) UpdateTimer(id string, o Job) error {
	return c.post("/api/timers/"+id, &o)
}
