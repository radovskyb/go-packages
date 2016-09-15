package main

import (
	"fmt"
	"log"
	"mime"
)

func main() {
	// .go is not yet associated with any current mime types
	fmt.Println(".go mime type:", mime.TypeByExtension(".go"))

	// Set to associate the mime type text/plain with the extension .go
	// which generally does not have an associated mime type
	//
	// AddExtensionType sets the MIME type associated with
	// the extension ext to typ. The extension should begin with
	// a leading dot, as in ".html".
	err := mime.AddExtensionType(".go", "text/plain")
	if err != nil {
		log.Fatalln(err)
	}

	// .go is now associated with the text/plain mime type
	fmt.Println(".go mime type:", mime.TypeByExtension(".go"))
}
