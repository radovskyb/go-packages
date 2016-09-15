package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// ResolveTCPAddr parses addr as a TCP address of the form "host:port"
	// or "[ipv6-host%zone]:port" and resolves a pair of domain name and
	// port name on the network net, which must be "tcp", "tcp4" or
	// "tcp6".  A literal address or host name for IPv6 must be enclosed
	// in square brackets, as in "[::1]:80", "[ipv6-host]:http" or
	// "[ipv6-host%zone]:80".
	tcpaddr, err := net.ResolveTCPAddr("tcp6", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(tcpaddr.Network())
	fmt.Println(tcpaddr.String())
	fmt.Println(tcpaddr.IP)
	fmt.Println(tcpaddr.Port)
	fmt.Println(tcpaddr.Zone)
}
