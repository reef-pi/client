package cluster

type Inlet struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Pin       int    `json:"pin"`
	Equipment string `json:"equipment"`
	Reverse   bool   `json:"reverse"`
	Driver    string `json:"driver"`
}

func (c *client) Inlets() ([]Inlet, error) {
	var inlets []Inlet
	return inlets, c.get("/api/inlets", &inlets)
}

func (c *client) Inlet(id string) (Inlet, error) {
	var inlet Inlet
	return inlet, c.get("/api/inlets/"+id, &inlet)
}

func (c *client) CreateInlet(o Inlet) error {
	return c.put("/api/inlets", &o)
}

func (c *client) DeleteInlet(id string) error {
	return c.delete("/api/inlets/" + id)
}

func (c *client) UpdateInlet(id string, o Inlet) error {
	return c.post("/api/inlets/"+id, &o)
}
