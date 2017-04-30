package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptrace"
)

func main() {
	// Package httptrace provides mechanisms to trace the events within
	// HTTP client requests.
	//
	// Create a new basic GET request.
	req, err := http.NewRequest("GET", "http://betsee.com.au", nil)
	if err != nil {
		log.Fatalln(err)
	}

	// Set up the tracing hooks to be used for the request.
	//
	// ClientTrace is a set of hooks to run at various stages of an outgoing
	// HTTP request. Any particular hook may be nil. Functions may be
	// called concurrently from different goroutines and some may be called
	// after the request has completed or failed.
	//
	// ClientTrace currently traces a single HTTP request & response
	// during a single round trip and has no hooks that span a series
	// of redirected requests.
	trace := &httptrace.ClientTrace{
		// GotConn is called after a successful connection is
		// obtained. There is no hook for failure to obtain a
		// connection; instead, use the error from
		// Transport.RoundTrip.
		GotConn: func(connInfo httptrace.GotConnInfo) {
			fmt.Printf("Got Conn: %+v\n", connInfo)
		},
		// DNSDone is called when a DNS lookup ends.
		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			fmt.Printf("DNS Info: %+v\n", dnsInfo)
		},
		// ConnectDone is called when a new connection's Dial
		// completes. The provided err indicates whether the
		// connection completedly successfully.
		// If net.Dialer.DualStack ("Happy Eyeballs") support is
		// enabled, this may be called multiple times.
		ConnectDone: func(network, addr string, err error) {
			fmt.Printf(
				"Connection Dial Complete: %s, %s, %+v\n",
				network, addr, err,
			)
		},
	}

	// Set the request's context so it will use the provided trace
	// hooks set in trace.
	//
	// WithClientTrace returns a new context based on the provided parent
	// ctx. HTTP client requests made with the returned context will use
	// the provided trace hooks, in addition to any previous hooks
	// registered with ctx. Any hooks defined in the provided trace will
	// be called first.
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))

	// Execute the actual request.
	_, err = http.DefaultTransport.RoundTrip(req)
	if err != nil {
		log.Fatal(err)
	}
}
