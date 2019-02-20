package client

import (
	"fmt"
	"testing"
)

func TestTimers(t *testing.T) {
	c := signIn(t)
	o := Job{
		Minute:    "*/2",
		Day:       "*",
		Hour:      "*",
		Second:    "1",
		Name:      "client-test-timer",
		Type:      "equipment",
		Equipment: UpdateEquipment{},
		Enable:    true,
	}

	eqs, eerr := c.ListEquipment()
	if eerr != nil {
		t.Error(eerr)
	}
	for _, eq := range eqs {
		o.Equipment.ID = eq.ID
		break
	}
	if err := c.CreateTimer(o); err != nil {
		t.Error(err)
	}
	timers, err := c.Timers()
	if err != nil {
		t.Error(err)
	}
	for _, o := range timers {
		fmt.Println(o.Name)
	}
	o.Name = "client-test-timer-updated"
	if err := c.UpdateTimer("1", o); err != nil {
		t.Error(err)
	}
	if err := c.CreateTimer(o); err != nil {
		t.Error(err)
	}
	if err := c.DeleteTimer("1"); err != nil {
		t.Error(err)
	}
	c.SignOut()
}
