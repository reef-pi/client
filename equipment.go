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
