package cluster

import (
	"time"
)

type ATO struct {
	ID             string        `json:"id"`
	Inlet          string        `json:"inlet"`
	Pump           string        `json:"pump"`
	Period         time.Duration `json:"period"`
	Control        bool          `json:"control"`
	Enable         bool          `json:"enable"`
	Notify         Notify        `json:"notify"`
	Name           string        `json:"name"`
	DisableOnAlert bool          `json:"disable_on_alert"`
}

func (c *client) ATOs() ([]ATO, error) {
	var atos []ATO
	return atos, c.get("/api/atos", &atos)
}

func (c *client) ATO(id string) (ATO, error) {
	var ato ATO
	return ato, c.get("/api/atos/"+id, &ato)
}

func (c *client) CreateATO(o ATO) error {
	return c.put("/api/atos", &o)
}

func (c *client) DeleteATO(id string) error {
	return c.delete("/api/atos/" + id)
}

func (c *client) UpdateATO(id string, o ATO) error {
	return c.post("/api/atos/"+id, &o)
}
