package main

import (
	"bytes"
	"fmt"
	"log"
)

func main() {
	// Create a byte slice
	b := []byte{'a', 'b'}

	// Create a new bytes buffer with b as its initial content
	var buf = bytes.NewBuffer(b)

	// Print the unread bytes from buffer buf before
	// reading any bytes from buffer buf
	//
	// Bytes returns a slice of the contents of the unread portion
	// of the buffer; len(b.Bytes()) == b.Len(). If the caller
	// changes the contents of the returned slice, the contents
	// of the buffer will change provided there are no
	// intervening method calls on the Buffer.
	fmt.Println(string(buf.Bytes())) // ab

	// Read one byte from buffer buf
	br, err := buf.ReadByte()
	if err != nil {
		log.Fatalln(err)
	}

	// Print the byte that was read from buffer buf as a string
	fmt.Println("Byte read from buffer buf:", string(br))

	// Now once again print the unread bytes from buffer buf
	fmt.Println(string(buf.Bytes())) // b
}
