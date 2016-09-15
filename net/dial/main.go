package main

import (
	"log"
	"net"
)

func main() {
	// Dial to localhost:8080 over tcp and return a connection to it if successful
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}

	// Write to the connection
	conn.Write([]byte("Hello, World!\n"))
}
