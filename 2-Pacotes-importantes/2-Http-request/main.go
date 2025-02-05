// INFO: 2.2-Chamada http
package main

import (
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(string(content))

	resp.Body.Close()
}
