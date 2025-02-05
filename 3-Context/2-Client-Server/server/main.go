// INFO: 3.2-Context utilizando server HTTP
package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Starting request")
	defer log.Println("Request finished")

	select {
	case <-time.After(5 * time.Second):
		log.Println("Request succesfully processed")
		w.Write([]byte("Request succesfully processed\n"))
    case <-ctx.Done():
		log.Println("Request cancelled")
        http.Error(w, "Request cancelled", http.StatusRequestTimeout)
	}
}
