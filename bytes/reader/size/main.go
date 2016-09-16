package main

import (
	"bytes"
	"fmt"
	"log"
)

func main() {
	// Create a new bytes reader with some contents
	r := bytes.NewReader([]byte{'a', 'b', 'c'})

	// Print r's size
	//
	// Size returns the original length of the underlying byte slice.
	// Size is the number of bytes available for reading via ReadAt.
	// The returned value is always the same and is not affected by
	// calls to any other method.
	fmt.Println(r.Size()) // 3

	// Read one byte from r
	_, err := r.ReadByte()
	if err != nil {
		log.Fatalln(err)
	}

	// Once again print the size
	fmt.Println(r.Size()) // 3
}
