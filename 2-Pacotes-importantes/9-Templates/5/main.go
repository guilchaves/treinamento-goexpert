// INFO: 2.14-Compondo templates
package main

import (
	"html/template"
	"net/http"
	"os"
)

type Class struct {
	Name     string
	Duration int
}

type Classes []Class

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	t := template.Must(template.New("content.html").ParseFiles(templates...))
	err := t.Execute(os.Stdout, Classes{
		{"Go", 40},
		{"Java", 70},
		{"Rust", 50},
	})
	if err != nil {
		panic(err)
	}
	http.ListenAndServe(":8000", nil)
}
