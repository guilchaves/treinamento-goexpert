// INFO: 8-Funcoes
package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(sum(50, 2))

	value, err := split(50, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(value)
}

func sum(a, b int) (int, bool) {
	if a+b >= 50 {
		return a + b, true
	}
	return a + b, false
}

func split(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Não é possível dividir por zero")
	}

	return a / b, nil
}
