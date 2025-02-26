// INFO: 2.15-Mapeando funcoes nos templates
package main

import (
	"html/template"
	"net/http"
	"os"
	"strings"
)

type Class struct {
	Name     string
	Duration int
}

type Classes []Class

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	t := template.New("content.html")
	t.Funcs(template.FuncMap{"ToUpper": ToUpper})
	t = template.Must(t.ParseFiles(templates...))

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
