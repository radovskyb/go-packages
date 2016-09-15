package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// LookupAddr performs a reverse lookup for the given address, returning
	// a list of names mapping to that address.
	//
	// Return a list of names mapping to the given address (127.0.0.1)
	names, err := net.LookupAddr("127.0.0.1")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(names) // [localhost]
}
