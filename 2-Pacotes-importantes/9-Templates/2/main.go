// INFO: 2.11-Templates.Must
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
    t := template.Must(template.New("ClassTemplate").Parse("Curso de {{.Name}} com duração de {{.Duration}} horas\n"))
    err := t.Execute(os.Stdout, class)
    if err != nil {
        panic(err)
    }

}
