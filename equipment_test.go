package cluster

import (
	"fmt"
	"testing"
)

func TestEquipment(t *testing.T) {
	c := siginIn(t)
	o := Equipment{
		Name: "client-test-equipment",
	}
	l := Outlet{
		Name:   "client-test-outlet",
		Driver: "rpi",
		Pin:    16,
	}
	if err := c.CreateOutlet(l); err != nil {
		t.Fatal(err)
	}
	outlets, oerr := c.Outlets()
	if oerr != nil {
		t.Error(oerr)
	}
	for _, l := range outlets {
		if l.Name == "client-test-outlet" {
			o.Outlet = l.ID
			break
		}
	}
	if err := c.CreateEquipment(o); err != nil {
		t.Error(err)
	}
	equipment, err := c.ListEquipment()
	if err != nil {
		t.Error(err)
	}
	for _, o := range equipment {
		fmt.Println(o.Name)
	}
	o.Name = "Bar"
	if err := c.UpdateEquipment("1", o); err != nil {
		t.Error(err)
	}
	if err := c.DeleteEquipment("1"); err != nil {
		t.Error(err)
	}
}
