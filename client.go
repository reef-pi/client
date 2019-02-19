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

func (c *client) get(p string, v interface{}) error {
	resp, err := c.c.Get(c.u + p)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("Failed http request. Status:%d, Error:%s", resp.StatusCode, string(body))
	}
	if v == nil {
		return nil
	}
	err = json.NewDecoder(resp.Body).Decode(v)
	return err
}

func (c *client) post(p string, v interface{}) error {
	buf := new(bytes.Buffer)
	if v != nil {
		if err := json.NewEncoder(buf).Encode(v); err != nil {
			return err
		}
	}
	req, err := http.NewRequest("POST", c.u+p, buf)
	if err != nil {
		return err
	}
	resp, err := c.c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("Failed http request. Status:%d, Error:%s", resp.StatusCode, string(body))
	}
	return nil
}
func (c *client) put(p string, v interface{}) error {
	buf := new(bytes.Buffer)
	if v != nil {
		if err := json.NewEncoder(buf).Encode(v); err != nil {
			return err
		}
	}
	req, err := http.NewRequest("PUT", c.u+p, buf)
	if err != nil {
		return err
	}
	resp, err := c.c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("Failed http request. Status:%d, Error:%s", resp.StatusCode, string(body))
	}
	return nil
}

func (c *client) delete(p string) error {
	buf := new(bytes.Buffer)
	req, err := http.NewRequest("DELETE", c.u+p, buf)
	if err != nil {
		return err
	}
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
