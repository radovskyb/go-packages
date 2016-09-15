package main

import (
	"fmt"
	"net"
)

func main() {
	// JoinHostPort combines host and port into a network address of the
	// form "host:port" or, if host contains a colon or a percent sign, "[host]:port".
	//
	// Join the host localhost and port 9000 into the
	// network address of the following form host:port
	addr := net.JoinHostPort("localhost", "9000")

	fmt.Println(addr)
}
