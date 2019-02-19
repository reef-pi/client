package cluster

import (
	"encoding/json"
)

type Channel struct {
	Name     string  `json:"name"`
	Min      int     `json:"min"`
	StartMin int     `json:"start_min"`
	Max      int     `json:"max"`
	Reverse  bool    `json:"reverse"`
	Pin      int     `json:"pin"`
	Color    string  `json:"color"`
	Profile  Profile `json:"profile"`
}
type Light struct {
	ID       string          `json:"id"`
	Name     string          `json:"name"`
	Channels map[int]Channel `json:"channels"`
	Jack     string          `json:"jack"`
}

type Profile struct {
	Type   string          `json:"type"`
	Config json.RawMessage `json:"config"`
}

type AutoConfig struct {
	Values []int `json:"values"` // 12 ticks after every 2 hours
}

type FixedConfig struct {
	Value int `json:"value"`
}

type DiurnalConfig struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

func (c *client) Lights() ([]Light, error) {
	var lights []Light
	return lights, c.get("/api/lights", &lights)
}

func (c *client) Light(id string) (Light, error) {
	var light Light
	return light, c.get("/api/lights/"+id, &light)
}

func (c *client) CreateLight(o Light) error {
	return c.put("/api/lights", &o)
}

func (c *client) DeleteLight(id string) error {
	return c.delete("/api/lights/" + id)
}

func (c *client) UpdateLight(id string, o Light) error {
	return c.post("/api/lights/"+id, &o)
}
