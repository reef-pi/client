package cluster

import (
	"fmt"
	"testing"
)

func siginIn(t *testing.T) *client {
	//u := "http://garage.local"
	u := "http://localhost:8080"
	c, err := NewClient(u)
	if err != nil {
		t.Fatal(err)
	}
	if err := c.SignIn("reef-pi", "reef-pi"); err != nil {
		t.Fatal(err)
	}
	return c
}

func TestAdmin(t *testing.T) {
	c := siginIn(t)
	if _, err := c.Capabilities(); err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	info, err := c.Info()
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println(info.Name)
	if err := c.Reload(); err != nil {
		t.Error(err)
	}
	eqs, err := c.ListEquipment()
	if err != nil {
		t.Error(err)
	}
	for _, eq := range eqs {
		fmt.Println(eq.Name)
	}
}
