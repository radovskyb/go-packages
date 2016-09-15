package main

import (
	"fmt"
	"net/http"
)

func main() {
	// CanonicalHeaderKey returns the canonical format of the
	// header key s.  The canonicalization converts the first
	// letter and any letter following a hyphen to upper case;
	// the rest are converted to lowercase.  For example, the
	// canonical key for "accept-encoding" is "Accept-Encoding".
	// If s contains a space or invalid header field bytes, it is
	// returned without modifications.
	chk := http.CanonicalHeaderKey("accept-encoding")

	// Print out the Canonical Header Key for the string `accept-encoding`
	fmt.Println(chk) // Accept-Encoding
}
