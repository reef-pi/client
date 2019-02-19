package cluster

import (
	"time"
)

type TC struct {
	ID         string        `json:"id"`
	Name       string        `json:"name"`
	Max        float64       `json:"max"`
	Min        float64       `json:"min"`
	Heater     string        `json:"heater"`
	Cooler     string        `json:"cooler"`
	Period     time.Duration `json:"period"`
	Control    bool          `json:"control"`
	Enable     bool          `json:"enable"`
	Notify     Notify        `json:"notify"`
	Sensor     string        `json:"sensor"`
	Fahrenheit bool          `json:"fahrenheit"`
	ChartMin   float64       `json:"chart_min"`
	ChartMax   float64       `json:"chart_max"`
}

type Notify struct {
	Enable bool    `json:"enable"`
	Max    float64 `json:"max"`
	Min    float64 `json:"min"`
}

func (c *client) TCs() ([]TC, error) {
	var tcs []TC
	return tcs, c.get("/api/tcs", &tcs)
}

func (c *client) TC(id string) (TC, error) {
	var tc TC
	return tc, c.get("/api/tcs/"+id, &tc)
}

func (c *client) CreateTC(o TC) error {
	return c.put("/api/tcs", &o)
}

func (c *client) DeleteTC(id string) error {
	return c.delete("/api/tcs/" + id)
}

func (c *client) UpdateTC(id string, o TC) error {
	return c.post("/api/tcs/"+id, &o)
}
