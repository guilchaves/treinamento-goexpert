//INFO: 13-Interfaces
package main

import "fmt"

type Address struct {
	Street string
	Number int
	City   string
	State  string
}

type Person interface {
	Deactivate()
}

type Firm struct {
	Name string
}

func (f Firm) Deactivate() {}

type Customer struct {
	Name   string
	Age    int
	Active bool
	Address
}

func (c Customer) Deactivate() {
	c.Active = false
	fmt.Printf("O cliente %s foi desativado\n", c.Name)
}

func Deactivation(p Person) {
	p.Deactivate()
}

func main() {
	guilherme := Customer{
		Name:   "Guilherme",
		Age:    32,
		Active: true,
	}

	myFirm := Firm{}

	Deactivation(guilherme)
	Deactivation(myFirm)
}
