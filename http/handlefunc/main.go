package main

import (
	"io"
	"log"
	"net/http"
)

func HelloWorld(w http.ResponseWriter, req *http.Request) {
	_, err := io.WriteString(w, "hello, world!\n")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	// HandleFunc registers the handler function for the given pattern
	// in the DefaultServeMux.
	// The documentation for ServeMux explains how patterns are matched.
	http.HandleFunc("/", HelloWorld)

	log.Fatal(http.ListenAndServe(":9000", nil))
}
