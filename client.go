package cluster

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"time"
)

func main() {
	fmt.Println("vim-go")
}

type client struct {
	u string
	c *http.Client
}

type credential struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

func NewClient(u string) (*client, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	c := &http.Client{
		Timeout: time.Second * 10,
		Jar:     jar,
	}
	return &client{
		u: u,
		c: c,
	}, nil
}

func (c *client) Login(u, p string) error {
	cred := credential{
		User:     u,
		Password: p,
	}
	br, err := json.Marshal(&cred)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", c.u+"/auth/signin", bytes.NewBuffer(br))
	resp, err := c.c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		if err != nil {
			body, _ := ioutil.ReadAll(resp.Body)
			return fmt.Errorf("Failed http request. Status:%d, Error:%s", resp.StatusCode, string(body))
		}
	}
	return nil
}

type Equipment struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Outlet string `json:"outlet"`
	On     bool   `json:"on"`
}

func (c *client) ListEquipment() ([]Equipment, error) {
	var eqs []Equipment
	resp, err := c.c.Get(c.u + "/api/equipment")
	if err != nil {
		return eqs, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		if err != nil {
			body, _ := ioutil.ReadAll(resp.Body)
			return eqs, fmt.Errorf("Failed http request. Status:%d, Error:%s", resp.StatusCode, string(body))
		}
	}

	return eqs, json.NewDecoder(resp.Body).Decode(&eqs)
}
