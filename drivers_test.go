package cluster

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDrivers(t *testing.T) {
	c := signIn(t)
	config := `{"address":64,"frequency":800}`
	o := Driver{
		Name:   "client-test-driver",
		Type:   "pca9685",
		Config: json.RawMessage(config),
	}
	if err := c.CreateDriver(o); err != nil {
		t.Error(err)
	}
	ds, err := c.Drivers()
	if err != nil {
		t.Error(err)
	}
	for _, o := range ds {
		fmt.Println(o.Name)
	}
	o.Name = "client-test-driver-updated"
	if err := c.UpdateDriver("1", o); err != nil {
		t.Error(err)
	}
	if err := c.CreateDriver(o); err != nil {
		t.Error(err)
	}
	if err := c.DeleteDriver("1"); err != nil {
		t.Error(err)
	}
	c.SignOut()
}
