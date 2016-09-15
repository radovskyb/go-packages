package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	// Create a new uint64 called varint
	var varint int64 = 12353123123123871

	// Make a []byte with the size of the size of bytes that
	// binary.Write would encode varint into
	buf := make([]byte, binary.Size(varint))

	// Encode varint and write the encoded value into buf
	//
	// PutVarint encodes an int64 into buf and returns the number
	// of bytes written.
	// If the buffer is too small, PutVarint will panic.
	n := binary.PutVarint(buf, varint)

	// Print out how many bytes were encoded and written to buf
	fmt.Println("binary.PutVarint wrote", n, "bytes")

	// Print the encoded varint stored in buf
	fmt.Println(buf)
}
