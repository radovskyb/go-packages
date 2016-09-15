package main

import (
	"fmt"
	"mime"
)

func main() {
	// FormatMediaType serializes mediatype t and the parameters
	// param as a media type conforming to RFC 2045 and RFC 2616.
	// The type and parameter names are written in lower-case.
	// When any of the arguments result in a standard violation then
	// FormatMediaType returns the empty string.
	s := mime.FormatMediaType("image/svg+xml", map[string]string{"svg": "\u0001"})
	fmt.Printf("%#v\n", s)
}
