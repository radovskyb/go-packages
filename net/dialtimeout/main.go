package main

import (
	"log"
	"net"
	"time"
)

func main() {
	// Dial to betsee.com.au:8080 and if the connection isn't successful
	// within a second then time out and disconnect
	// (port 8080 doesn't exist at betsee so it will timeout)
	//
	// DialTimeout acts like Dial but takes a timeout. The timeout includes
	// name resolution, if required.
	_, err := net.DialTimeout("tcp", "betsee.com.au:8080", time.Second)
	if err != nil {
		log.Fatalln(err)
	}
}
