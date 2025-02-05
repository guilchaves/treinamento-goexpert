// INFO: 2.13-Templates com webserver
package main

import (
	"net/http"
	"text/template"
)

type Class struct {
	Name     string
	Duration int
}

type Classes []Class

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("template.html").ParseFiles("template.html"))
		err := t.Execute(w, Classes{
			{"Go", 40},
			{"Java", 70},
			{"Rust", 50},
		})
		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":8000", nil)
}
