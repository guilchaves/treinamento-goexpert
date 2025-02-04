//INFO: 7-Maps
package main

import "fmt"

func main() {

	salarios := map[string]int{"Guilherme": 1000, "José": 1500, "Maria": 2000}

	// Different ways to create a map
	//sal := make(map[string]int)
	//sal1 := map[string]int{}

	for key, value := range salarios {
		fmt.Printf("O salario de %s é %d\n", key, value)
	}

	for _, salario := range salarios {
		fmt.Println(salario)
	}

}
