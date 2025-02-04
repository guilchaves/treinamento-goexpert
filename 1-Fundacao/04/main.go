//INFO: 4-Importando fmt e tipagem
package main

import "fmt"

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
    fmt.Printf("O tipo de E é %T\n", e) //float64
    fmt.Printf("O tipo de F é %T\n", f) //main.ID
}
