package cluster

import (
	"fmt"
	"testing"
)

func TestOutlets(t *testing.T) {
	c := siginIn(t)
	o := Outlet{
		Name:   "foo",
		Pin:    23,
		Driver: "rpi",
	}
	if err := c.CreateOutlet(o); err != nil {
		t.Error(err)
	}
	outlets, err := c.Outlets()
	if err != nil {
		t.Error(err)
	}
	for _, o := range outlets {
		fmt.Println(o.Name)
	}
	o.Name = "Bar"
	if err := c.UpdateOutlet("1", o); err != nil {
		t.Error(err)
	}
	if err := c.CreateOutlet(o); err != nil {
		t.Error(err)
	}
	if err := c.DeleteOutlet("1"); err != nil {
		t.Error(err)
	}
}
