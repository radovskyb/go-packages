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
	// HandleFunc registers the handler function for the given pattern
	// in the DefaultServeMux.
	// The documentation for ServeMux explains how patterns are matched.
	http.HandleFunc("/", HelloWorld)

	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
