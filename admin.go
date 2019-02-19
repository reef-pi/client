package cluster

/*
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)
*/

type Capabilities struct {
	DevMode       bool `json:"dev_mode"`
	Dashboard     bool `json:"dashboard"`
	HealthCheck   bool `json:"health_check"`
	Equipment     bool `json:"equipment"`
	Timers        bool `json:"timers"`
	Lighting      bool `json:"lighting"`
	Temperature   bool `json:"temperature"`
	ATO           bool `json:"ato"`
	Camera        bool `json:"camera"`
	Doser         bool `json:"doser"`
	Ph            bool `json:"ph"`
	Macro         bool `json:"macro"`
	Configuration bool `json:"configuration"`
}

func (c *client) SignIn(u, p string) error {
	cred := credential{
		User:     u,
		Password: p,
	}
	return c.post("/auth/signin", &cred)
}

func (c *client) SignOut() error {
	return c.get("/auth/signout", nil)
}

func (c *client) Capabilities() (Capabilities, error) {
	var caps Capabilities
	return caps, c.get("/api/capabilities", &caps)
}

func (c *client) PowerOff() error {
	return c.post("/api/admin/poweroff", nil)
}

func (c *client) Reboot() error {
	return c.post("/api/admin/reboot", nil)
}

func (c *client) Reload() error {
	return c.post("/api/admin/reload", nil)
}

type Info struct {
	Name           string `json:"name"`
	IP             string `json:"ip"`
	CurrentTime    string `json:"current_time"`
	Uptime         string `json:"uptime"`
	CPUTemperature string `json:"cpu_temperature"`
	Version        string `json:"version"`
}

func (c *client) Info() (Info, error) {
	var i Info
	return i, c.get("/api/info", &i)
}
