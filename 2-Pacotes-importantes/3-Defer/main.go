// INFO: 2.3-Defer
package main

import (
	"io"
	"net/http"
)

func main() {
	// Defer is LIFO ordered (works as a stack)
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(string(content))
}
