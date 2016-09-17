package main

import (
	"fmt"
	"log"
	"net/http"
)

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	// Redirect replies to the request with a redirect to url,
	// which may be a path relative to the request path.
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(w, "Home Page<br /><a href=\"/redirect\">Go To Redirect Page</a>")
}

func main() {
	http.HandleFunc("/redirect", RedirectHandler)
	http.HandleFunc("/", HomeHandler)
	log.Fatal(http.ListenAndServe(":9000", nil))
}
