package cluster

type Pump struct {
	ID       string         `json:"id"`
	Name     string         `json:"name"`
	Jack     string         `json:"jack"`
	Pin      int            `json:"pin"`
	Regiment DosingRegiment `json:"regiment"`
}
type DosingRegiment struct {
	Enable   bool     `json:"enable"`
	Schedule Schedule `json:"schedule"`
	Duration float64  `json:"duration"`
	Speed    float64  `json:"speed"`
}

type CalibrationDetails struct {
	Speed    float64 `json:"speed"`
	Duration float64 `json:"duration"`
}

type Schedule struct {
	Day    string `json:"day"`
	Hour   string `json:"hour"`
	Minute string `json:"minute"`
	Second string `json:"second"`
}

func (c *client) Dosers() ([]Pump, error) {
	var dosers []Pump
	return dosers, c.get("/api/dosers", &dosers)
}

func (c *client) Doser(id string) (Pump, error) {
	var doser Pump
	return doser, c.get("/api/dosers/"+id, &doser)
}

func (c *client) CreateDoser(o Pump) error {
	return c.put("/api/dosers", &o)
}

func (c *client) DeleteDoser(id string) error {
	return c.delete("/api/dosers/" + id)
}

func (c *client) UpdateDoser(id string, o Pump) error {
	return c.post("/api/dosers/"+id, &o)
}
