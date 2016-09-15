package main

import (
	"fmt"
	"mime"
)

func main() {
	// Print out all of the known file extensions associated with
	// the mime type text/plain which are found from this current
	// computers list of mime types. Therefore, the list can be
	// different on different systems that has a different list
	//
	// ExtensionsByType returns the extensions known to be associated with the MIME
	// type typ. The returned extensions will each begin with a leading dot, as in
	// ".html". When typ has no associated extensions, ExtensionsByType returns an
	// nil slice.
	fmt.Println(mime.ExtensionsByType("text/plain"))
}
