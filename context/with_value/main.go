package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func printUsernameHandler(w http.ResponseWriter, r *http.Request) {
	// Value returns the value associated with this context for key, or nil
	// if no value is associated with key. Successive calls to Value with
	// the same key returns the same result.
	username := r.Context().Value("username").(string)
	if strings.TrimSpace(username) == "" {
		http.Error(w, "username is blank", http.StatusBadRequest)
		return
	}
	if username == "admin" {
		fmt.Fprintln(w, "Welcome Administrator")
		return
	}
	fmt.Fprintf(w, "Hello, %s\n", strings.ToUpper(username))
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		splitPath := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
		if len(splitPath) < 2 || splitPath[0] != "username" {
			http.Error(w, "username is not set", http.StatusBadRequest)
			return
		}
		// Create a shallow copy of r with it's context changed to a new context that
		// is the same as r's current context but with an added value of a username.
		//
		// The new request object is then passed to the printUsernameHandler handler function.
		//
		// WithValue returns a copy of parent in which the value associated with key is
		// val.
		//
		// Use context Values only for request-scoped data that transits processes and
		// APIs, not for passing optional parameters to functions.
		//
		// The provided key must be comparable.
		newReq := r.WithContext(context.WithValue(
			r.Context(), "username", splitPath[1],
		))
		printUsernameHandler(w, newReq)
	})
	log.Fatal(http.ListenAndServe(":9000", nil))
}
