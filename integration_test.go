package client

import (
	"testing"
)

func signIn(t *testing.T) *client {
	//u := "http://garage.local"
	u := "http://localhost:8080"
	c, err := New(u)
	if err != nil {
		t.Fatal(err)
	}
	if err := c.SignIn("reef-pi", "reef-pi"); err != nil {
		t.Fatal(err)
	}
	return c
}
