package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// Create a new transport
	t := &http.Transport{}

	// NewFileTransport returns a new RoundTripper, serving the provided
	// FileSystem. The returned RoundTripper ignores the URL host in its
	// incoming requests, as well as most other properties of the
	// request.
	//
	// The typical use case for NewFileTransport is to register the "file"
	// protocol with a Transport, as in:
	//
	//   t := &http.Transport{}
	//   t.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))
	//   c := &http.Client{Transport: t}
	//   res, err := c.Get("file:///etc/passwd")
	//   ...
	t.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))

	// Create a new client that uses transport t from above as it's transport
	c := &http.Client{Transport: t}

	// Get the /etc/ folder using a get request
	resp, err := c.Get("file:///etc/")
	if err != nil {
		log.Fatalln(err)
	}

	// Write out the response to os.Stdout
	if err := resp.Write(os.Stdout); err != nil {
		log.Fatalln(err)
	}
}
