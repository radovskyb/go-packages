package main

import (
	"fmt"
	"strings"
)

func main() {
	// Create a new strings reader for the string below
	//
	// NewReader returns a new Reader reading from s.
	// It is similar to bytes.NewBufferString but more
	// efficient and read-only.
	rdr := strings.NewReader("Hello, World!")

	// Len returns the number of bytes of the unread portion of the
	// string.
	fmt.Println(rdr.Len()) // 13
}
