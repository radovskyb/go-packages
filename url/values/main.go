package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	v := url.Values{}
	v.Set("name", "Benjamin Radovsky")
	v.Add("friend", "Friend 1")
	v.Add("friend", "Friend 2")

	// Print out the encoded url values.
	//
	// Encode encodes the values into ``URL encoded'' form
	// ("bar=baz&foo=quux") sorted by key.
	encoded := v.Encode()

	// Create a base url.
	base, err := url.Parse("http://localhost/")
	if err != nil {
		log.Fatalln(err)
	}

	// Print out the base url before adding the encoded
	// values as it's RawQuery.
	fmt.Println("Before:", base)

	// Add the encoded valued to the base url.
	base.RawQuery = encoded

	// Now print out the url which will show the entire url
	// with the encoded query appended to it.
	fmt.Println("After:", base)
}
