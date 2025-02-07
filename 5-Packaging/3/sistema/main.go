//INFO: 5.5-Trabalhando com go mod replace
package main

/* To run localy -> Set the package to the local dir:
go mod edit -replace github.com/guilchaves/treinamento-goexpert/5-Packaging/3/math=../math */
import "github.com/guilchaves/treinamento-goexpert/5-Packaging/3/math"

func main() {
	m := math.NewMath(1, 2)
	println(m.Add())
}
