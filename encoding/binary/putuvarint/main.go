package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	// Create a new uint64 called uvarint
	var uvarint uint64 = 12353123123123871

	// Make a []byte with the size of the size of bytes that
	// binary.Write would encode uvarint into
	buf := make([]byte, binary.Size(uvarint))

	// Encode uvarint and write the encoded value into buf
	//
	// PutUvarint encodes a uint64 into buf and returns the number
	// of bytes written.
	// If the buffer is too small, PutUvarint will panic.
	n := binary.PutUvarint(buf, uvarint)

	// Print out how many bytes were encoded and written to buf
	fmt.Println("binary.PutUvarint wrote", n, "byte/s")

	// Print the encoded uvarint stored in buf
	fmt.Println(buf)
}
