package main

import (
	"bytes"
	"fmt"
	"log"
)

func main() {
	// Create a new bytes reader with some contents
	r := bytes.NewReader([]byte{'a', 'b', 'c'})

	// Create a new byte slice of 3 bytes in size
	b := make([]byte, 3)

	// Now read from reader r into b with an offset of 1
	_, err := r.ReadAt(b, 1)
	if err != nil {
		log.Fatalln(err)
	}

	// Print the contents of b that were read in from r
	// Will only print `bc` since the offset was 1
	fmt.Println(string(b))
}
