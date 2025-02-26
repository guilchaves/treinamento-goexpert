// INFO: 9-Funcoes variadicas
package main

import (
	"fmt"
)

func main() {
	fmt.Println((sum(1, 2, 4, 55, 8, 3, 23, 88, 4, 2, 68, 2, 8, 19)))
}

// Variadic function
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
