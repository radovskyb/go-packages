package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	// Parse a url.
	u, err := url.Parse("https://example.com/foo%2fbar")
	if err != nil {
		log.Fatal(err)
	}

	// Print out the url's request uri.
	//
	// RequestURI returns the encoded path?query or opaque?query
	// string that would be used in an HTTP request for u.
	fmt.Println(u.RequestURI())
}
