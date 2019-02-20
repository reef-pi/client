package client

import (
	"fmt"
	"testing"
)

func TestPhProbes(t *testing.T) {
	c := signIn(t)
	o := Probe{
		Name:        "client-test-probe",
		AnalogInput: "2",
		Period:      10,
	}
	if err := c.CreatePhProbe(o); err != nil {
		t.Error(err)
	}
	probes, err := c.PhProbes()
	if err != nil {
		t.Error(err)
	}
	for _, o := range probes {
		fmt.Println(o.Name)
	}
	o.Name = "client-test-probe-updated"
	if err := c.UpdatePhProbe("1", o); err != nil {
		t.Error(err)
	}
	if err := c.CreatePhProbe(o); err != nil {
		t.Error(err)
	}
	if err := c.DeletePhProbe("1"); err != nil {
		t.Error(err)
	}
	c.SignOut()
}
