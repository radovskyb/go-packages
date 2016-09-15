package main

import (
	"fmt"
	"net"
	"net/http"
	"sync"
)

const addr = "localhost:9000"

type homeHandler struct {
	sync.Mutex
	count int
}

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Home</h1>")
	var count int
	h.Lock()
	h.count++
	count = h.count
	h.Unlock()

	fmt.Fprintf(w, "Visitor count: %d.", count)
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
	mux.HandleFunc("/about", aboutHandler)

	server := &http.Server{
		Addr:      addr,
		Handler:   mux,
		ConnState: printState,
	}

	server.ListenAndServe()
}
