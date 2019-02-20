package cluster

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

func (c *client) ControlEquipment(id string, on bool) error {
	o := make(map[string]bool)
	o["on"] = on
	return c.post("/api/equipment/"+id+"/control", &o)
}
