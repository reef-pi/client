package cluster

import (
	"time"
)

type Probe struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Enable      bool          `json:"enable"`
	Period      time.Duration `json:"period"`
	AnalogInput string        `json:"analog_input"`
	Control     bool          `json:"control"`
	Notify      Notify        `json:"notify"`
}

func (c *client) PhProbes() ([]Probe, error) {
	var probes []Probe
	return probes, c.get("/api/phprobes", &probes)
}

func (c *client) PhProbe(id string) (Probe, error) {
	var outlet Probe
	return outlet, c.get("/api/phprobes/"+id, &outlet)
}

func (c *client) CreatePhProbe(o Probe) error {
	return c.put("/api/phprobes", &o)
}

func (c *client) DeletePhProbe(id string) error {
	return c.delete("/api/phprobes/" + id)
}

func (c *client) UpdatePhProbe(id string, o Probe) error {
	return c.post("/api/phprobes/"+id, &o)
}
