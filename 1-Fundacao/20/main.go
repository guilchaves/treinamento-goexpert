// INFO: 20-Pacotes e módulos
package main

import (
    "fmt"
    "curso-go/matematica"
)

func main() {
	sum := matematica.Sum(10, 20)
	fmt.Printf("Resultado: %v\n", sum)
}
