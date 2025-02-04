//INFO: 3-Criacao de tipos
package main

const a = "This value never changes"

type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "Guilherme"
	e float64 = 1.2
    f ID = 1
)

func main() {
    a := "x"
	println(a) // x
	println(f) // x
}
