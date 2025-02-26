// INFO: 11-Iniciando com structs
package main

import "fmt"

type Address struct {
	Street string
	Number int
	City   string
	State  string
}

type Customer struct {
	Name   string
	Age    int
	Active bool
	Address
}

func main() {
	guilherme := Customer{
		Name:   "Guilherme",
		Age:    32,
		Active: true,
	}

	guilherme.Active = false
	guilherme.Street = "Rua de Mogi"
	guilherme.Number = 123
	guilherme.City = "Mogi das Cruzes"
	guilherme.State = "SP"

	fmt.Printf("Name: %s, Age: %d, Active: %t\n", guilherme.Name, guilherme.Age, guilherme.Active)
}
