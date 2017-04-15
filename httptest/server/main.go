package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"
)

type homeHandler struct {
	mu    sync.Mutex
	count int
}

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Home</h1>")

	h.mu.Lock()
	h.count++
	fmt.Fprintf(w, "Visitor count: %d.\n", h.count)
	h.mu.Unlock()
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>About</h1>")
}

func printState(conn net.Conn, state http.ConnState) {
	fmt.Println(state)
}

func main() {
	mux := http.NewServeMux()
	homeHandler := &homeHandler{}
	mux.Handle("/", homeHandler)
	// Ignore any /favicon.ico GET requests from the browser to prevent
	// double incrementing of the homeHandler count.
	mux.HandleFunc("/favicon.ico", http.NotFound)
	mux.HandleFunc("/about", aboutHandler)

	server := &http.Server{
		Addr:      "localhost:9000",
		Handler:   mux,
		ConnState: printState,
	}

	log.Fatal(server.ListenAndServe())
}
