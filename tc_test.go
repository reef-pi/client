package client

import (
	"fmt"
	"testing"
)

func TestTCs(t *testing.T) {
	c := signIn(t)
	o := TC{
		Name:     "client-test-tc",
		Period:   10,
		Enable:   true,
		Sensor:   "sdad",
		ChartMin: 72,
		ChartMax: 82,
	}
	if err := c.CreateTC(o); err != nil {
		t.Error(err)
	}
	tcs, err := c.TCs()
	if err != nil {
		t.Error(err)
	}
	for _, o := range tcs {
		fmt.Println(o.Name)
	}
	o.Name = "client-test-tc-updated"
	if err := c.UpdateTC("1", o); err != nil {
		t.Error(err)
	}
	if err := c.CreateTC(o); err != nil {
		t.Error(err)
	}
	if err := c.DeleteTC("1"); err != nil {
		t.Error(err)
	}
	c.SignOut()
}
