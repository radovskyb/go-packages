package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

// PrintState simple listens on the server for changed states
// and just prints any changes to os.Stdout
func PrintState(conn net.Conn, state http.ConnState) {
	fmt.Println(state)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Homepage</h1>")
}

func main() {
	// Create a new serve mux to use with the custom &http.Server{}
	mux := http.NewServeMux()

	// Register a handler for the path "/" on the serve mux `mux`
	mux.HandleFunc("/", homeHandler)

	// Create a new &http.Server{}
	//
	// A Server defines parameters for running an HTTP server.
	// The zero value for Server is a valid configuration.
	server := &http.Server{
		Addr:         ":9000",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		ConnState:    PrintState,
	}

	// Listen and serve using server which will listen on port 9000
	// in it's own go routine so the second listener and server.Serve()
	// call can run below, therefore making it as though there are 2
	// servers running, one on port 9000, and one on port 9001
	//
	// ListenAndServe listens on the TCP network address srv.Addr and then
	// calls Serve to handle requests on incoming connections.  If
	// srv.Addr is blank, ":http" is used.
	go server.ListenAndServe()

	// Create a tcp listener on port 9001
	ln, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalln(err)
	}

	// Separately serve the server `server` on the tcp listener `ln`
	//
	// server.Serve will close the listener when it's done with it
	// so there's no need to call defer ln.Close() above.
	log.Fatal(server.Serve(ln))
}
