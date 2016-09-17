package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	h := http.FileServer(http.Dir("/tmp"))

	// Timeout duration
	td := time.Second

	// TimeoutHandler returns a Handler that runs h with the given time limit.
	//
	// The new Handler calls h.ServeHTTP to handle each request, but if a
	// call runs for longer than its time limit, the handler responds with
	// a 503 Service Unavailable error and the given message in its body.
	// (If msg is empty, a suitable default message will be sent.)
	// After such a timeout, writes by h to its ResponseWriter will return
	// ErrHandlerTimeout.
	log.Fatal(http.ListenAndServe(":9000", http.TimeoutHandler(h, td, "File scanning timed out!")))
}
