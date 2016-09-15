package main

import (
	"fmt"
	"net/http"
)

func main() {
	// ParseHTTPVersion parses a HTTP version string.
	// "HTTP/1.0" returns (1, 0, true).
	major, minor, ok := http.ParseHTTPVersion("HTTP/1.0")

	// Print out the major and minor HTTP version string as well
	// as if the string was a correct HTTP Version string (ok)
	fmt.Println(major, minor, ok)
}
