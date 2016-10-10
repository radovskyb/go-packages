package main

import (
	"fmt"
	"log"
	"syscall"
)

func main() {
	// Create a new socket
	sockfd, err := syscall.Socket(syscall.AF_INET6, syscall.SOCK_STREAM, 0)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		// Close the socket
		if err := syscall.Close(sockfd); err != nil {
			log.Fatalln(err)
		}
	}()

	// Bind the socket to port 9000
	sockaddr := &syscall.SockaddrInet6{Port: 9000}

	// Connect to the socket
	if err := syscall.Connect(sockfd, sockaddr); err != nil {
		log.Fatalln(err)
	}

	// Create a new buffer to store data read from the socket
	buf := make([]byte, 256)

	// Read data from the socket connection into the buffer
	n, err := syscall.Read(sockfd, buf)
	if err != nil {
		log.Fatalln(err)
	}

	// Print out the data read from the socket connection
	fmt.Printf("Read: %s\n", buf[0:n])
}
