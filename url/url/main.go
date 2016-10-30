package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	// Parse a url.
	//
	// Parse parses rawurl into a URL structure.
	// The rawurl may be relative or absolute.
	u, err := url.Parse("https://example.com/foo%2fbar#im_a_fragment")
	if err != nil {
		log.Fatal(err)
	}

	// Print out the url's path.
	// Prints /foo/bar even though the url was encoded as foo%2fbar.
	fmt.Println(u.Path)

	// Print out the url's raw path.
	fmt.Println(u.RawPath)

	// Print out the url's fragment.
	fmt.Println(u.Fragment)

	// Print out the url as a string.
	fmt.Println(u.String())

	// Print a blank line before the next url's prints.
	fmt.Println()

	// Parse a url.
	u, err = url.Parse("http://localhost:9000/search?name=Benjamin Radovsky")
	if err != nil {
		log.Fatalln(err)
	}

	// Print out the url's scheme.
	fmt.Println(u.Scheme)

	// Print out the url's host.
	fmt.Println(u.Host)

	// Print out the url's path.
	fmt.Println(u.Path)

	// Print out the url's raw query.
	fmt.Println(u.RawQuery)

	// Encode and then print out the encoded query.
	//
	// Query parses RawQuery and returns the corresponding values.
	//
	// Encode encodes the values into ``URL encoded'' form
	// ("bar=baz&foo=quux") sorted by key.
	fmt.Println(u.Query().Encode())

	// Print out the whole url.
	fmt.Println(u)
}
