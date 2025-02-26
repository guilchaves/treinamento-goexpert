// INFO: 18-Type assertions
package main

import "fmt"

func main() {
	var myVar interface{} = "Guilherme Chaves"

	println(myVar.(string))
	res, ok := myVar.(int)
	fmt.Printf("O valor de res é %v e o resultado de ok é %v\n", res, ok)
}
