package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// Create an integer channel msg
var msg = make(chan int)

func handler(w http.ResponseWriter, r *http.Request) {
	// Set the correct headers to allow the connection to
	// be kept alive and the content type to allow continuous
	// message sending through channel msg
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// Create a multi writer that will write to both os.Stdout
	// and w at the same time
	mw := io.MultiWriter(os.Stdout, w)

	// Check if w connection is capable of using the interface
	// http.CloseNotifier and if it isn't then return an http.Error
	if _, ok := w.(http.CloseNotifier); !ok {
		http.Error(w, "cannot stream", http.StatusInternalServerError)
		return
	}

	// Start a go routine that sends an integer along the msg channel
	// once every second
	go func() {
		for i := 0; i < 3; i++ {
			msg <- i
			time.Sleep(time.Second)
		}
	}()

LOOP:
	for {
		select {
		case msg := <-msg:
			fmt.Fprintln(mw, "Received", msg)
			w.(http.Flusher).Flush()
		case <-time.After(time.Second * 3):
			break LOOP
		}
	}

	// Call CloseNotify() which will notify both os.Stdout and
	// w when the connection is closed
	//
	// The CloseNotifier interface is implemented by ResponseWriters which
	// allow detecting when the underlying connection has gone away.
	//
	// This mechanism can be used to cancel long operations on the server
	// if the client has disconnected before the response is ready.
	w.(http.CloseNotifier).CloseNotify()
	fmt.Fprintln(mw, "Disconnected")
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":9000", nil))
}
