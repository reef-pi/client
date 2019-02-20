package cluster

import (
	"fmt"
	"testing"
)

func TestInlets(t *testing.T) {
	c := signIn(t)
	o := Inlet{
		Name:   "baz",
		Pin:    21,
		Driver: "rpi",
	}
	if err := c.CreateInlet(o); err != nil {
		t.Error(err)
	}
	inlets, err := c.Inlets()
	if err != nil {
		t.Error(err)
	}
	for _, o := range inlets {
		fmt.Println(o.Name)
	}
	o.Name = "client-test-inlet-updated"
	if err := c.UpdateInlet("1", o); err != nil {
		t.Error(err)
	}
	if err := c.CreateInlet(o); err != nil {
		t.Error(err)
	}
	if err := c.DeleteInlet("1"); err != nil {
		t.Error(err)
	}
	c.SignOut()

}
