package main

import (
	"io"
	"log"
	"net/http"
)

func HelloWorld(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
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
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
