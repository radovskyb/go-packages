package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// LookupMX returns the DNS MX records for the given domain name sorted by preference.
	mxs, err := net.LookupMX("google.com")
	if err != nil {
		log.Fatalln(err)
	}

	for _, mx := range mxs {
		fmt.Println(mx)
	}
}
