package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
)

func Client(c net.Conn) {
	// Create a new bytes buffer
	var buf bytes.Buffer
	// Copy anything written to c out to bytes buffer buf
	_, err := io.Copy(&buf, c)
	if err != nil {
		log.Fatalln(err)
	}

	// Append to and then print the string inside of buf to stdout
	fmt.Print(buf.String() + " printed from client connection")
}

func Server(c net.Conn) {
	// Send a byte slice over the connection
	_, err := c.Write([]byte("Hello, world from server connection."))
	if err != nil {
		log.Fatalln(err)
	}
	// Close the connection to tell the pipe that writing is done
	if err := c.Close(); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	// Pipe creates a synchronous, in-memory, full duplex
	// network connection; both ends implement the Conn interface.
	// Reads on one end are matched with writes on the other,
	// copying data directly between the two; there is no internal
	// buffering.
	c1, c2 := net.Pipe()
	// Call Server(c1) which will write to it's connection
	go Server(c1)
	// Call Client(c2) which will read from it's connection and print to stdout
	Client(c2)
}
