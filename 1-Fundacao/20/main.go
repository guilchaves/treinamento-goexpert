// INFO: 20-Pacotes e módulos
package main

import (
	"curso-go/matematica"
	"fmt"
)

func main() {
	sum := matematica.Sum(10, 20)
	fmt.Printf("Resultado: %v\n", sum)
}
