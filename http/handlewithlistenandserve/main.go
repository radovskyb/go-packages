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
	// Handle registers the handler for the given pattern
	// in the DefaultServeMux.
	// The documentation for ServeMux explains how patterns are matched.
	http.Handle("/", http.HandlerFunc(HelloWorld))

	// ListenAndServe listens on the TCP network address addr
	// and then calls Serve with handler to handle requests
	// on incoming connections.  Handler is typically nil,
	// in which case the DefaultServeMux is used.
	log.Fatal(http.ListenAndServe(":9000", nil))
}
