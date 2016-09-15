package main

import (
	"fmt"
	"net/textproto"
)

func main() {
	// Return the canonical format of the mime
	// header key `accept-Encoding`
	//
	// CanonicalMIMEHeaderKey returns the canonical format of the
	// MIME header key s.  The canonicalization converts the first
	// letter and any letter following a hyphen to upper case;
	// the rest are converted to lowercase.  For example, the
	// canonical key for "accept-encoding" is "Accept-Encoding".
	// MIME header keys are assumed to be ASCII only.
	// If s contains a space or invalid header field bytes, it is
	// returned without modifications.
	str := textproto.CanonicalMIMEHeaderKey("accept-encoding")
	fmt.Println(str) // Accept-Encoding
}
