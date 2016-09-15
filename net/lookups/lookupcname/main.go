package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// LookupCNAME returns the canonical DNS host for the given name.
	// Callers that do not care about the canonical name can call
	// LookupHost or LookupIP directly; both take care of resolving
	// the canonical name as part of the lookup.
	cname, err := net.LookupCNAME("ftp.betsee.com.au")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(cname) // betsee.com.au
}
