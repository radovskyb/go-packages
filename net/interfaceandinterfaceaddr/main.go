package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// InterfaceAddrs returns a list of the system's network interface addresses.
	addresses, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatalln(err)
	}

	// Print out a list of the system's network interface addresses
	for _, address := range addresses {
		fmt.Println(address)
	}

	fmt.Println()

	// Interfaces returns a list of the system's network interfaces.
	ifaces, err := net.Interfaces()
	if err != nil {
		log.Fatalln(err)
	}

	// Print out a list of the system's network interfaces
	for _, iface := range ifaces {
		fmt.Println(iface)
	}
}
