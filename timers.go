package cluster

import (
	"time"
)

type Reminder struct {
	Title   string `json:"title"`
	Message string `json:"message"`
}

type UpdateEquipment struct {
	Revert   bool          `json:"revert"`
	ID       string        `json:"id"`
	On       bool          `json:"on"`
	Duration time.Duration `json:"duration"`
}
type Job struct {
	ID        string          `json:"id"`
	Minute    string          `json:"minute"`
	Day       string          `json:"day"`
	Hour      string          `json:"hour"`
	Second    string          `json:"second"`
	Name      string          `json:"name"`
	Type      string          `json:"type"`
	Reminder  Reminder        `json:"reminder"`
	Equipment UpdateEquipment `json:"equipment"`
	Enable    bool            `json:"enable"`
}

func (c *client) Timers() ([]Job, error) {
	var timers []Job
	return timers, c.get("/api/timers", &timers)
}

func (c *client) Timer(id string) (Job, error) {
	var timer Job
	return timer, c.get("/api/timers/"+id, &timer)
}

func (c *client) CreateTimer(o Job) error {
	return c.put("/api/timers", &o)
}

func (c *client) DeleteTimer(id string) error {
	return c.delete("/api/timers/" + id)
}

func (c *client) UpdateTimer(id string, o Job) error {
	return c.post("/api/timers/"+id, &o)
}
