// INFO: 2.10-Iniciando com templates
package main

import (
	"os"
	"text/template"
)

type Class struct {
	Name     string
	Duration int
}

func main() {
	class := Class{"Go", 40}
	temp := template.New("ClassTemplate")
	temp, _ = temp.Parse("Curso de {{.Name}} com duração de {{.Duration}} horas\n")
    err := temp.Execute(os.Stdout, class)
    if err != nil {
        panic(err)
    }

}
