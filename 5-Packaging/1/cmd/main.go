// INFO: 5.1-Intoduzindo go mod | 5.2-Acessando pacotes criados | 5.3-Exportação de objetos
package main

import (
	"fmt"
	"github.com/guilchaves/treinamento-goexpert/5-Packaging/1/math"
)

func main() {
	m := math.Math{A: 1, B: 2}
	fmt.Println(m.Add())
}
