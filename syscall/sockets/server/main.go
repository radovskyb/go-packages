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

	if err := syscall.Bind(sockfd, sockaddr); err != nil {
		log.Fatalln(err)
	}

	// Allow connections to be accepted on the socket
	if err := syscall.Listen(sockfd, syscall.SOMAXCONN); err != nil {
		log.Fatalln(err)
	}

	// Loop forever, accepting new connections on the socket
	for {
		// Accept a new connection on the socket
		newsockfd, _, err := syscall.Accept(sockfd)
		if err != nil {
			log.Fatalln(err)
		}

		// Handle the new socket connection in it's own goroutine
		go handleConn(newsockfd)
	}
}

func handleConn(sockfd int) {
	defer syscall.Close(sockfd) // Close the socket connection

	// Write to the new socket connection
	_, err := syscall.Write(sockfd, []byte("Hello, World!"))
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Sent data to socket connection.")
}
