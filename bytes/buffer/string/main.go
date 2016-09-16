package main

import (
	"bytes"
	"fmt"
	"log"
)

func main() {
	// Create a new bytes buffer with 5 bytes as it's
	// starting contents (abcde)
	buf := bytes.NewBuffer([]byte{'a', 'b', 'c', 'd', 'e'})

	// Read one byte from buf
	_, err := buf.ReadByte() // Reads `a`
	if err != nil {
		log.Fatalln(err)
	}

	// Print the remaining contents using buf.String()
	//
	// String returns the contents of the unread portion of the buffer
	// as a string. If the Buffer is a nil pointer, it returns "<nil>".
	fmt.Println(buf.String()) // (bcde)
}
