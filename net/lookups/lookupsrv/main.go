package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// LookupSRV tries to resolve an SRV query of the given service,
	// protocol, and domain name.  The proto is "tcp" or "udp".
	// The returned records are sorted by priority and randomized
	// by weight within a priority.
	//
	// LookupSRV constructs the DNS name to look up following RFC 2782.
	// That is, it looks up _service._proto.name.  To accommodate services
	// publishing SRV records under non-standard names, if both service
	// and proto are empty strings, LookupSRV looks up name directly.
	cname, addrs, err := net.LookupSRV("xmpp-server", "tcp", "google.com")
	if err != nil {
		log.Fatalln(err)
	}

	// Print out the cname
	fmt.Println(cname)

	// Print out the addrs
	for _, addr := range addrs {
		fmt.Println(addr)
	}
}
