package client

func (c *Client) ListEquipment() ([]Equipment, error) {
	var eqs []Equipment
	return eqs, c.get("/api/equipment", &eqs)
}

func (c *Client) Equipment(id string) (Equipment, error) {
	var ato Equipment
	return ato, c.get("/api/equipment/"+id, &ato)
}

func (c *Client) CreateEquipment(o Equipment) error {
	return c.put("/api/equipment", &o)
}

func (c *Client) DeleteEquipment(id string) error {
	return c.delete("/api/equipment/" + id)
}

func (c *Client) UpdateEquipment(id string, o Equipment) error {
	return c.post("/api/equipment/"+id, &o)
}

func (c *Client) ControlEquipment(id string, on bool) error {
	o := make(map[string]bool)
	o["on"] = on
	return c.post("/api/equipment/"+id+"/control", &o)
}
