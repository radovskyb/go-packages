package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// SplitHostPort splits a network address of the form "host:port",
	// "[host]:port" or "[ipv6-host%zone]:port" into host or
	// ipv6-host%zone and port.  A literal address or host name for IPv6
	// must be enclosed in square brackets, as in "[::1]:80",
	// "[ipv6-host]:http" or "[ipv6-host%zone]:80".
	host, port, err := net.SplitHostPort("localhost:9000")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(host)
	fmt.Println(port)
}
