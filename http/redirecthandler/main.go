package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", handler)

	// RedirectHandler returns a request handler that redirects
	// each request it receives to the given url using the given
	// status code.
	http.Handle("/redirect", http.RedirectHandler("/",
		http.StatusTemporaryRedirect))
	http.ListenAndServe(":9000", nil)
}
