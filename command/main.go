package main

import (
	"fmt"
	"github.com/reef-pi/cluster"
)

const u = "http://garage.local"

func main() {
	c, err := cluster.NewClient(u)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	if err := c.Login("reef-pi", "reef-pi"); err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	eqs, err := c.ListEquipment()
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	for _, eq := range eqs {
		fmt.Println(eq.Name)
	}
}
