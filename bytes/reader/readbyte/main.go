package main

import (
	"bytes"
	"fmt"
	"log"
)

func main() {
	// Create a new bytes reader with some contents
	r := bytes.NewReader([]byte{'a', 'b', 'c'})

	// Read one byte from r
	b, err := r.ReadByte()
	if err != nil {
		log.Fatalln(err)
	}

	// Print the byte that was read from r.ReadByte()
	fmt.Println(string(b))
}
