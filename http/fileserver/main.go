package main

import (
	"log"
	"net/http"
)

func main() {
	// Create a simple static webserver using http.FileServer
	// for the files folder from in this directory
	//
	// FileServer returns a handler that serves HTTP requests
	// with the contents of the file system rooted at root.
	//
	// To use the operating system's file system implementation,
	// use http.Dir:
	//
	//     http.Handle("/", http.FileServer(http.Dir("/tmp")))
	//
	// As a special case, the returned file server redirects any request
	// ending in "/index.html" to the same path, without the final
	// "index.html".
	log.Fatal(http.ListenAndServe(":9000", http.FileServer(http.Dir("./files"))))
}
