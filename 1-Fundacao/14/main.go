// INFO: 14-Ponteiros
package main

func main() {
	// Memoria -> Endereco -> Valor
	a := 10
	var pointer *int = &a
	*pointer = 20

	println(a)

	b := &a
	println(*b)

}
