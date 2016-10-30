package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	// username and password must be in plain text (see below).
	rawURL := "http://admin:pass@localhost:9000/search/?q=Benjamin#BenjiFragment"

	// Parse the query from rawURL.
	//
	// ParseQuery parses the URL-encoded query string and returns
	// a map listing the values specified for each key.
	// ParseQuery always returns a non-nil map containing all the
	// valid query parameters found; err describes the first decoding error
	// encountered, if any.
	val, err := url.ParseQuery(rawURL)
	if err != nil {
		log.Fatalln(err)
	}

	// Print out the raw query key and value.
	for k, v := range val {
		fmt.Printf("%v:\n%v\n", k, v)
	}
}
