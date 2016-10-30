package main

import (
	"fmt"
	"net/url"
)

func main() {
	// Escape the query and store it in query string variable.
	//
	// QueryEscape escapes the string so it can be safely placed
	// inside a URL query.
	escaped := url.QueryEscape("name=Benjamin Radovsky&age=24")

	// Print out the escaped query.
	fmt.Println(escaped)
}
