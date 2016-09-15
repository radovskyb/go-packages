package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// LookupHost looks up the given host using the local resolver. It
	// returns an array of that host's addresses.
	addresses, err := net.LookupHost("localhost")
	if err != nil {
		log.Fatalln(err)
	}

	// Print out an array of the host 'localhost's addresses (::1, 127.0.0.1)
	for _, address := range addresses {
		fmt.Println(address)
	}
}
