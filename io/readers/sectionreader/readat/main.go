package main

import (
	"fmt"
	"io"
	"log"
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

	// Create a new byte slice to read into
	buf := make([]byte, 6)

	// Read from sr with an offset of 1
	_, err := sr.ReadAt(buf, 1) // Returns: orld!
	if err != nil && err != io.EOF {
		log.Fatalln(err)
	}

	// Print buf to os.Stdout
	fmt.Println(string(buf)) // orld!
}
