package cluster

import (
	"testing"
)

func TestAdmin(t *testing.T) {
	c := signIn(t)
	if _, err := c.Capabilities(); err != nil {
		t.Error(err)
	}
	if _, err := c.Info(); err != nil {
		t.Error(err)
	}
	if err := c.Reload(); err != nil {
		t.Error(err)
	}
	if err := c.PowerOff(); err != nil {
		t.Error(err)
	}
	if err := c.Reboot(); err != nil {
		t.Error(err)
	}
	if err := c.SignOut(); err != nil {
		t.Error(err)
	}
}
