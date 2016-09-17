package main

import (
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Hijack the response writer's connection.
	conn, _, err := w.(http.Hijacker).Hijack()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write over the hijacked, now raw tcp connection.
	_, err = conn.Write([]byte("Server has hijacked this connection."))
	if err != nil {
		panic(err)
	}

	// Close the connection when we're done.
	if err := conn.Close(); err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":9000", nil))
}
