package cluster

import (
	"fmt"
	"testing"
)

func TestATOs(t *testing.T) {
	c := siginIn(t)
	o := ATO{
		Inlet:   "1",
		Pump:    "1",
		Period:  10,
		Control: true,
		Enable:  true,
		Notify:  Notify{},
		Name:    "client-test-ato",
	}
	if err := c.CreateATO(o); err != nil {
		t.Error(err)
	}
	atos, err := c.ATOs()
	if err != nil {
		t.Error(err)
	}
	for _, o := range atos {
		fmt.Println(o.Name)
	}
	o.Name = "Bar"
	if err := c.UpdateATO("1", o); err != nil {
		t.Error(err)
	}
	if err := c.CreateATO(o); err != nil {
		t.Error(err)
	}
	if err := c.DeleteATO("1"); err != nil {
		t.Error(err)
	}
}
