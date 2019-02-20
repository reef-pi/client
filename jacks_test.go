package cluster

import (
	"fmt"
	"testing"
)

func TestJacks(t *testing.T) {
	c := signIn(t)
	o := Jack{
		Name:   "test-jack",
		Pins:   []int{0},
		Driver: "rpi",
	}
	if err := c.CreateJack(o); err != nil {
		t.Error(err)
	}
	jacks, err := c.Jacks()
	if err != nil {
		t.Error(err)
	}
	for _, o := range jacks {
		fmt.Println(o.Name)
	}
	o.Name = "client-test-jack-updated"
	if err := c.UpdateJack("1", o); err != nil {
		t.Error(err)
	}
	if err := c.CreateJack(o); err != nil {
		t.Error(err)
	}
	if err := c.DeleteJack("1"); err != nil {
		t.Error(err)
	}
	c.SignOut()
}
