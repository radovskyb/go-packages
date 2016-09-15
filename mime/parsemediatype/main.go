package main

import (
	"fmt"
	"log"
	"mime"
)

func main() {
	// ParseMediaType parses a media type value and any optional
	// parameters, per RFC 1521.  Media types are the values in
	// Content-Type and Content-Disposition headers (RFC 2183).
	// On success, ParseMediaType returns the media type converted
	// to lowercase and trimmed of white space and a non-nil map.
	// The returned map, params, maps from the lowercase
	// attribute to the attribute value with its case preserved.
	mType, parameters, err := mime.ParseMediaType("text/plain; charset=utf-8; v=x")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Media type:\n%s\n", mType)

	if len(parameters) > 0 {
		fmt.Println("\nParameters:")
	}
	for param := range parameters {
		fmt.Printf("%v: %v\n", param, parameters[param])
	}
}
