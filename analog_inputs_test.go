package cluster

import (
	"fmt"
	"testing"
)

func TestAnalogInputs(t *testing.T) {
	c := signIn(t)
	o := AnalogInput{
		Name: "foo",
	}
	d := Driver{
		Name:   "baex",
		Type:   "ph-board",
		Config: []byte(`{"address":45}`),
	}
	if err := c.CreateDriver(d); err != nil {
		t.Error(err)
	}
	ds, err1 := c.Drivers()
	if err1 != nil {
		t.Error(err1)
	}
	for _, d := range ds {
		if d.Name == "baex" {
			o.Driver = d.ID
		}
	}
	if err := c.CreateAnalogInput(o); err != nil {
		t.Error(err)
	}
	as, err := c.AnalogInputs()
	if err != nil {
		t.Error(err)
	}
	for _, o := range as {
		fmt.Println(o.Name)
	}
	o.Name = "Bar"
	if err := c.UpdateAnalogInput("1", o); err != nil {
		t.Error(err)
	}
	if err := c.CreateAnalogInput(o); err != nil {
		t.Error(err)
	}
	if err := c.DeleteAnalogInput("1"); err != nil {
		t.Error(err)
	}
}
