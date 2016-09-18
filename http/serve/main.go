package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World using http.Serve() over a tcp listener")
}

func main() {
	ln, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		log.Fatalln(err)
	}

	http.HandleFunc("/", handler)

	// http.Serve will close the listener when it's done with it
	// so there's no need to call defer ln.Close() above.
	//
	// Serve accepts incoming HTTP connections on the listener l,
	// creating a new service goroutine for each.  The service goroutines
	// read requests and then call handler to reply to them.
	// Handler is typically nil, in which case the DefaultServeMux is used.
	log.Fatal(http.Serve(ln, nil))
}
