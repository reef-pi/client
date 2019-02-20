package client

import (
	"fmt"
	"testing"
)

func TestMacros(t *testing.T) {
	c := signIn(t)
	o := Macro{
		Name: "client-test-macro",
	}
	if err := c.CreateMacro(o); err != nil {
		t.Error(err)
	}
	outlets, err := c.Macros()
	if err != nil {
		t.Error(err)
	}
	for _, o := range outlets {
		fmt.Println(o.Name)
	}
	o.Name = "client-test-macro-updated"
	if err := c.UpdateMacro("1", o); err != nil {
		t.Error(err)
	}
	if err := c.CreateMacro(o); err != nil {
		t.Error(err)
	}
	if err := c.DeleteMacro("1"); err != nil {
		t.Error(err)
	}
	c.SignOut()
}
