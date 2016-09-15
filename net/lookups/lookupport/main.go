package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// Look up the port for the network type tcp and the service telnet
	//
	// LookupPort looks up the port for the given network and service.
	port, err := net.LookupPort("tcp", "telnet")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(port)
}
