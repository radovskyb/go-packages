package main

import (
	"bufio"
	"bytes"
	"log"
	"net/http"
	"os"
)

func main() {
	// Create a new bytes buffer containing a get request
	bytesRequest := bytes.NewBuffer([]byte("GET / HTTP/1.1\r\nheader:foo\r\n\r\n"))

	// Create a new buffered reader forthe bytes buffer containing a GET request
	reader := bufio.NewReader(bytesRequest)

	// Read and parse the request from the buffered reader
	//
	// ReadRequest reads and parses an incoming request from b.
	req, err := http.ReadRequest(reader)
	if err != nil {
		log.Fatalln(err)
	}

	// Write the request to os.Stdout
	if err := req.Write(os.Stdout); err != nil {
		log.Fatalln(err)
	}
}
