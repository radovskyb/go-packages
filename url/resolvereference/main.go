package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	// Parse a partial non-absolute url.
	u, err := url.Parse("../../../search?q=gophers")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u.IsAbs())

	// Parse an absolute base url.
	base, err := url.Parse("http://example.com/one/two/three/four/")
	if err != nil {
		log.Fatal(err)
	}

	// Print out the base's resolved reference passing it the first url.
	//
	// ResolveReference resolves a URI reference to an absolute URI from
	// an absolute base URI, per RFC 3986 Section 5.2.  The URI reference
	// may be relative or absolute.  ResolveReference always returns a new
	// URL instance, even if the returned URL is identical to either the
	// base or reference. If ref is an absolute URL, then ResolveReference
	// ignores base and returns a copy of ref.
	fmt.Println(base.ResolveReference(u))
}
