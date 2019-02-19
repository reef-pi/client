package cluster

type Equipment struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Outlet string `json:"outlet"`
	On     bool   `json:"on"`
}

func (c *client) ListEquipment() ([]Equipment, error) {
	var eqs []Equipment
	return eqs, c.get("/api/equipment", &eqs)
}

func (c *client) Equipment(id string) (Equipment, error) {
	var ato Equipment
	return ato, c.get("/api/equipment/"+id, &ato)
}

func (c *client) CreateEquipment(o Equipment) error {
	return c.put("/api/equipment", &o)
}

func (c *client) DeleteEquipment(id string) error {
	return c.delete("/api/equipment/" + id)
}

func (c *client) UpdateEquipment(id string, o Equipment) error {
	return c.post("/api/equipment/"+id, &o)
}
