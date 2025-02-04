//INFO: 2-Declaracao e atribuicao
package main

const a = "This value never changes"

var (
	b bool    = true
	c int     = 10
	d string  = "Guilherme"
	e float64 = 1.2
)

func main() {
    a := "x"
	println(a) // x
	println(b)
	println(c)
	println(d)
	println(e)
}
