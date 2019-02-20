package client

func (c *client) Macros() ([]Macro, error) {
	var macros []Macro
	return macros, c.get("/api/macros", &macros)
}

func (c *client) Macro(id string) (Macro, error) {
	var m Macro
	return m, c.get("/api/macros/"+id, &m)
}

func (c *client) CreateMacro(o Macro) error {
	return c.put("/api/macros", &o)
}

func (c *client) DeleteMacro(id string) error {
	return c.delete("/api/macros/" + id)
}

func (c *client) UpdateMacro(id string, o Macro) error {
	return c.post("/api/macros/"+id, &o)
}

func (c *client) RunMacro(id string) error {
	return c.post("/api/macros/"+id+"/run", nil)
}
