// INFO: 17-Interfaces vazias
package main

import "fmt"

// Sempre buscar usar generics
func main() {

	var x interface{} = 10
	var y interface{} = "hello"

	showType(x)
	showType(y)
}

func showType(t interface{}) {
	fmt.Printf("O tipo da variável é %T e o valor é %v\n", t, t)
}
