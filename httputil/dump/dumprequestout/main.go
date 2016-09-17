package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
)

func main() {
	const body = "Go is a general-purpose language designed with systems programming in mind."

	req, err := http.NewRequest("PUT", "http://betsee.com.au", strings.NewReader(body))
	if err != nil {
		log.Fatalln(err)
	}

	// DumpRequestOut is like DumpRequest but includes
	// headers that the standard http.Transport adds,
	// such as User-Agent.
	dump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%q", dump)
}
