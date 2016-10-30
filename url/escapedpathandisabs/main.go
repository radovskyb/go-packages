package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	// Parse a url
	u, err := url.Parse("http://localhost:9000/search name?name=Benjamin Radovsky")
	if err != nil {
		log.Fatalln(err)
	}

	// Print out the url's non escaped path.
	fmt.Println(u.Path) // search name

	// Print the url's escaped path.
	//
	// EscapedPath returns the escaped form of u.Path.
	// In general there are multiple possible escaped forms of any path.
	// EscapedPath returns u.RawPath when it is a valid escaping of u.Path.
	// Otherwise EscapedPath ignores u.RawPath and computes an escaped
	// form on its own.
	// The String and RequestURI methods use EscapedPath to construct
	// their results.
	// In general, code should call EscapedPath instead of
	// reading u.RawPath directly.
	fmt.Println(u.EscapedPath()) // search%20name

	// Print out whether the url is absolute or not.
	fmt.Println(u.IsAbs())

	fmt.Println()

	// Parse another url.
	u, err = url.Parse("localhost?name=Benjamin")
	if err != nil {
		log.Fatalln(err)
	}

	// Print out whether the url is absolute or not.
	fmt.Println(u.IsAbs())
}
