package main

import (
	"log"
	"net/http"
)

func main() {
	// StripPrefix returns a handler that serves HTTP requests
	// by removing the given prefix from the request URL's Path
	// and invoking the handler h. StripPrefix handles a
	// request for a path that doesn't begin with prefix by
	// replying with an HTTP 404 not found error.
	http.Handle("/tmpfiles/", http.StripPrefix("/tmpfiles/", http.FileServer(http.Dir("/tmp"))))
	log.Fatal(http.ListenAndServe(":9000", nil))
}
