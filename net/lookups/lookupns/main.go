package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// LookupNS returns the DNS NS records for the given domain name.
	nss, err := net.LookupNS("betsee.com.au")
	if err != nil {
		log.Fatalln(err)
	}

	// Print out a list of name server records for betsee.com.au
	for _, ns := range nss {
		fmt.Println(ns)
	}
}
