package client

func (c *Client) Macros() ([]Macro, error) {
	var macros []Macro
	return macros, c.get("/api/macros", &macros)
}

func (c *Client) Macro(id string) (Macro, error) {
	var m Macro
	return m, c.get("/api/macros/"+id, &m)
}

func (c *Client) CreateMacro(o Macro) error {
	return c.put("/api/macros", &o)
}

func (c *Client) DeleteMacro(id string) error {
	return c.delete("/api/macros/" + id)
}

func (c *Client) UpdateMacro(id string, o Macro) error {
	return c.post("/api/macros/"+id, &o)
}

func (c *Client) RunMacro(id string) error {
	return c.post("/api/macros/"+id+"/run", nil)
}
