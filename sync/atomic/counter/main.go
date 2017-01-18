package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
)

// hitCounter is a middleware handler that increments it's
// counter value every time a new request is received.
type hitCounter struct {
	hits uint64
}

func (h *hitCounter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Increment the hit counter's internal counter value and print the
	// current number of hits that have occurred to the visitor.
	//
	// AddUint64 atomically adds delta to *addr and returns the new value.
	// To subtract a signed positive constant value c from x, do AddUint64(&x, ^uint64(c-1)).
	// In particular, to decrement x, do AddUint64(&x, ^uint64(0)).
	fmt.Fprintf(w, "visitor number %d\n", atomic.AddUint64(&h.hits, 1))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Index Page</h1>")
}

func main() {
	// Create a new hit counter.
	hc := new(hitCounter)

	// Load indexHandler when any page is requested.
	http.HandleFunc("/", indexHandler)

	// ListenAndServe at localhost:9000, serving all
	// requests through the hc handler.
	log.Fatal(http.ListenAndServe(":9000", hc))

	// To test hit counter.
	//
	// 1. Start application.
	// 2. Run the command `curl localhost:9000` in the
	//	  terminal a few times.
}
