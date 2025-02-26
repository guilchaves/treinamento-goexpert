// INFO: 2.12-Templates com arquivo externo
package main

import (
	"os"
	"text/template"
)

type Class struct {
	Name     string
	Duration int
}

type Classes []Class

func main() {
	t := template.Must(template.New("template.html").ParseFiles("template.html"))
	err := t.Execute(os.Stdout, Classes{
		{"Go", 40},
		{"Java", 70},
		{"Rust", 50},
	})
	if err != nil {
		panic(err)
	}
}
