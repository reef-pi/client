package client

import (
	"fmt"
	"testing"
)

func TestDosers(t *testing.T) {
	c := signIn(t)
	o := Pump{
		Name:     "client-test-doser",
		Jack:     "2",
		Pin:      1,
		Regiment: DosingRegiment{},
	}
	if err := c.CreateDoser(o); err != nil {
		t.Error(err)
	}
	dosers, err := c.Dosers()
	if err != nil {
		t.Error(err)
	}
	for _, o := range dosers {
		fmt.Println(o.Name)
	}
	o.Name = "client-test-doser-updated"
	if err := c.UpdateDoser("1", o); err != nil {
		t.Error(err)
	}
	if err := c.CreateDoser(o); err != nil {
		t.Error(err)
	}
	if err := c.DeleteDoser("1"); err != nil {
		t.Error(err)
	}
	c.SignOut()
}
