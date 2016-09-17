package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", handler)

	// NotFoundHandler returns a simple request handler
	// that replies to each request with a ``404 page not found'' reply.
	http.Handle("/notfound", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":9000", nil))
}
