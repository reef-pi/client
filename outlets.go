package cluster

type Outlet struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Pin       int    `json:"pin"`
	Equipment string `json:"equipment"`
	Reverse   bool   `json:"reverse"`
	Driver    string `json:"driver"`
}

func (c *client) Outlets() ([]Outlet, error) {
	var outlets []Outlet
	return outlets, c.get("/api/outlets", &outlets)
}

func (c *client) Outlet(id string) (Outlet, error) {
	var outlet Outlet
	return outlet, c.get("/api/outlets/"+id, &outlet)
}

func (c *client) CreateOutlet(o Outlet) error {
	return c.put("/api/outlets", &o)
}

func (c *client) DeleteOutlet(id string) error {
	return c.delete("/api/outlets/" + id)
}

func (c *client) UpdateOutlet(id string, o Outlet) error {
	return c.post("/api/outlets/"+id, &o)
}
