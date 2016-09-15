package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// Look up google.com.au's ip's
	// LookupIP looks up host using the local resolver. It returns an
	// array of that host's IPv4 and IPv6 addresses.
	ips, err := net.LookupIP("google.com.au")
	if err != nil {
		log.Fatalln(err)
	}

	// Print out the list of ip addresses
	for _, ip := range ips {
		fmt.Println(ip)
	}
}
