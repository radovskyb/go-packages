package main

import (
	"fmt"
	"log"
	"net/http"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("http.NotFound from NotFounderHandler on page", r.URL.Path)

	// NotFound replies to the request with an HTTP 404 not found error.
	http.NotFound(w, r)
}

func main() {
	// Handler NotFoundHandler will call http.NotFound which will
	// write a 404 error to the user's request
	http.HandleFunc("/", NotFoundHandler)

	// Now do the same thing directly using http.HandlerFunc(http.NotFound)
	// which simply wraps the above type of NotFoundHandler method in
	// a simple call rather than having to make a whole handler to do the job
	http.HandleFunc("/notfound", http.HandlerFunc(http.NotFound))

	log.Fatal(http.ListenAndServe(":9000", nil))
}
