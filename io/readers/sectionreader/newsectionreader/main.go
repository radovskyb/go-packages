package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	// Create a new string s
	s := "Hello, World!"

	// Create a new strings reader r to read from the string s
	r := strings.NewReader(s)

	// Create a new section reader that starts at the offset of
	// 7 bytes and reads 6 bytes
	//
	// NewSectionReader returns a SectionReader that reads from r
	// starting at offset off and stops with EOF after n bytes.
	sr := io.NewSectionReader(r, 7, 6)

	// Copy sr's contents to os.Stdout
	_, err := io.Copy(os.Stdout, sr) // Print: World!
	if err != nil {
		log.Fatalln(err)
	}
}
