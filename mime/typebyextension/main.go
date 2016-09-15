package main

import (
	"fmt"
	"mime"
	"path/filepath"
)

func main() {
	filenames := []string{
		"scripts.js",
		"index.html",
		"log.txt",
		"database.php",
		"server.go",
	}

	var ext, mType string

	for _, filename := range filenames {
		ext = filepath.Ext(filename)

		// TypeByExtension returns the MIME type associated with the file
		// extension ext.
		// The extension ext should begin with a leading dot, as in ".html".
		// When ext has no associated type, TypeByExtension returns "".
		//
		// Extensions are looked up first case-sensitively, then case-insensitively.
		//
		// The built-in table is small but on unix it is augmented by the local
		// system's mime.types file(s) if available under one or more of these
		// names:
		//
		//   /etc/mime.types
		//   /etc/apache2/mime.types
		//   /etc/apache/mime.types
		//
		// On Windows, MIME types are extracted from the registry.
		//
		// Text types have the charset parameter set to "utf-8" by default.
		mType = mime.TypeByExtension(ext)
		if mType != "" {
			fmt.Printf("%s (%s) has media type: %s\n", filename, ext, mType)
		} else {
			fmt.Printf("%s (%s) has an unknown media type\n", filename, ext)
		}
	}
}
