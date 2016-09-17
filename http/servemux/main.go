package main

import (
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" && r.URL.Path != "/home" {
		http.NotFound(w, r)
		return
	}

	_, err := w.Write([]byte("<h1>Home Page</h1><a href=\"\\about\">About Page</a>"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("<h1>About Page</h1><a href=\"\\home\">Home Page</a>"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ServHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("<h1>Get Page</h1><a href=\"\\home\">Home Page</a>"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	// Create a new serve mux called `mux`
	//
	// NewServeMux allocates and returns a new ServeMux.
	mux := http.NewServeMux()

	// Handle registers the handler for the given pattern.
	// If a handler already exists for pattern, Handle panics.
	mux.Handle("/", http.HandlerFunc(homeHandler))

	// HandleFunc registers the handler function for the given pattern.
	mux.HandleFunc("/about", aboutHandler)

	mux.HandleFunc("/handler", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")

		// Create a new handler
		//
		// Handler returns the handler to use for the given request,
		// consulting r.Method, r.Host, and r.URL.Path. It always returns
		// a non-nil handler. If the path is not in its canonical form, the
		// handler will be an internally-generated handler that redirects
		// to the canonical path.
		//
		// Handler also returns the registered pattern that matches the
		// request or, in the case of internally-generated redirects,
		// the pattern that will match after following the redirect.
		//
		// If there is no registered handler that applies to the request,
		// Handler returns a ``page not found'' handler and an empty pattern.
		h, _ := mux.Handler(r)

		// Create a new request to pass to the h.ServeHTTP() method
		req, err := http.NewRequest("GET", "/home", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Call h's ServeHTTP method and pass the current writer
		// and the newly created request req
		h.ServeHTTP(w, req)
	})

	// Listen and serve on port 9000 using the serve mux `mux`
	log.Fatal(http.ListenAndServe(":9000", nil))
}
