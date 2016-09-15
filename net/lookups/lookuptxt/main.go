package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// LookupTXT returns the DNS TXT records for the given domain name.
	txts, err := net.LookupTXT("betsee.com.au")
	if err != nil {
		log.Fatalln(err)
	}

	for _, txt := range txts {
		fmt.Println(txt)
	}
}
