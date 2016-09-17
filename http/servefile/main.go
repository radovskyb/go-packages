package main

import (
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// ServeFile replies to the request with the contents of the named
	// file or directory.
	//
	// As a special case, ServeFile redirects any request where r.URL.Path
	// ends in "/index.html" to the same path, without the final
	// "index.html". To avoid such redirects either modify the path or
	// use ServeContent.
	http.ServeFile(w, r, "hello.txt")
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":9000", nil))
}
