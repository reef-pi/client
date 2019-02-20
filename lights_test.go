package cluster

import (
	"fmt"
	"testing"
)

func TestLights(t *testing.T) {
	c := signIn(t)
	ch := Channel{
		Name: "test-client",
		Pin:  1,
		Profile: Profile{
			Type:   "fixed",
			Config: []byte(`{"value":23}`),
		},
	}
	chs := make(map[int]Channel)
	chs[1] = ch
	o := Light{
		Name:     "foo",
		Channels: chs,
	}
	j := Jack{
		Driver: "rpi",
		Name:   "client-test-TestLightsk",
		Pins:   []int{1},
	}
	if err := c.CreateJack(j); err != nil {
		t.Error(err)
	}
	jacks, jErr := c.Jacks()
	if jErr != nil {
		t.Error(jErr)
	}
	for _, jack := range jacks {
		if j.Name == jack.Name {
			o.Jack = jack.ID
		}
	}
	if err := c.CreateLight(o); err != nil {
		t.Error(err)
	}
	lights, err := c.Lights()
	if err != nil {
		t.Error(err)
	}
	for _, o := range lights {
		fmt.Println(o.Name)
	}
	o.Name = "client-test-light-updated"
	if err := c.UpdateLight("1", o); err != nil {
		t.Error(err)
	}
	if err := c.CreateLight(o); err != nil {
		t.Error(err)
	}
	if err := c.DeleteLight("1"); err != nil {
		t.Error(err)
	}
	c.SignOut()
}
