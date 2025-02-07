// INFO: 5.6-Usando workspaces
package main

/*
Create a workspace
> at root dir:
go work init ./math /.sistema

*/
import (
    "github.com/guilchaves/treinamento-goexpert/5-Packaging/4/math"
    "github.com/google/uuid"
)

func main() {
	m := math.Math{A: 1, B: 2}
	println(m.Add())
}
